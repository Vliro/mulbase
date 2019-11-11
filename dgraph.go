package mulbase

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc/encoding/gzip"
	"io/ioutil"
	"strconv"
	"sync"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type DB struct {
	d        *dgo.Dgraph
	ip       string
	port     int
	tls      bool
	endPoint string
	schema   schemaList
}

//NewTxn creates a new txn for interacting with database.
func (d *DB) NewTxn(readonly bool) *Txn {
	if d.schema == nil {
		panic("transaction without schema")
	}
	txn := new(Txn)
	if readonly {
		txn.txn = d.d.NewReadOnlyTxn()
	} else {
		txn.txn = d.d.NewTxn()
	}
	txn.sch = d.schema
	return txn
}

//Queries outside a Txn context.
//This is not intended for mutations.
func (d *DB) Query(ctx context.Context, q Query, obj interface{}) error {
	if q.Type() != QueryRegular {
		return Error(errInvalidType)
	}
	txn := d.NewTxn(true)
	err := txn.RunQuery(ctx, q, obj)
	//TODO: Can we do this for readonly?
	_ = txn.Discard(ctx)
	return err
}

func (d *DB) Mutate(ctx context.Context, q Query) error {
	if q.Type() != QuerySet || q.Type() != QueryDelete {
		return Error(errInvalidType)
	}
	txn := d.NewTxn(false)
	err := txn.RunQuery(ctx, q, nil)
	if err != nil {
		_ = txn.Discard(ctx)
	} else {
		_ = txn.Commit(ctx)
	}
	return err
}

//Sets the database schema. This is required for any query.
func (d *DB) SetSchema(sch map[string]Field) {
	if sch == nil {
		panic("nil schema map supplied")
	}
	d.schema = sch
}
//Simply performs the alter command.
func (d *DB) Alter(ctx context.Context, op *api.Operation) error {
	err := d.d.Alter(ctx, op)
	return err
}

type QueryType uint8

const (
	QueryRegular QueryType = iota
	QuerySet
	QueryDelete
)

type Query interface {
	//Process the query type in order to send to the database.
	Process(schemaList) ([]byte, map[string]string, error)
	//What type of query is this? Mutation(set/delete), regular query?
	Type() QueryType
}

//tlsPaths has length 5.
//First parameter is the root CA.
//Second is the client crt, third is the client key.
//Fourth is the node crt, fifth is the node key.
func Init(dip string, dport int, tls bool, tlsPath ...string) *DB {
	if dport < 1000 {
		panic("graphinit: invalid dgraph port number")
	}
	db := connect(dip, dport, tls, tlsPath)
	return db
	//initSchema()
}

//Takes a connection create of the form http://ip:port/
func connect(ip string, port int, dotls bool, paths []string) *DB {
	//TODO: allow multiple dgraph clusters.
	var conn *grpc.ClientConn
	var err error
	if dotls {
		//TODO: Review this code as it might change how TLS works with dgraph.
		rootCAs := x509.NewCertPool()
		cCerts, err := tls.LoadX509KeyPair(paths[3], paths[4])
		certs, err := ioutil.ReadFile(paths[0])
		rootCAs.AppendCertsFromPEM(certs)
		conf := &tls.Config{}
		conf.RootCAs = rootCAs
		conf.Certificates = append(conf.Certificates, cCerts)
		c := credentials.NewTLS(conf)
		conn, err = grpc.Dial(ip+":"+strconv.Itoa(port), grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)), grpc.WithTransportCredentials(c))
		if err != nil {
			panic(err)
		}
	} else {
		conn, err = grpc.Dial(ip+":"+strconv.Itoa(port), grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)), grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
	}
	//TODO: For multiple DGraph servers append multiple connections here.
	var c = dgo.NewDgraphClient(api.NewDgraphClient(conn))
	db := &DB{
		d:        c,
		ip:       ip,
		port:     port,
		tls:      dotls,
		endPoint: ip + ":" + strconv.Itoa(port) + "/graphql",
	}
	return db
}

//Txn is a non thread-safe API for interacting with the database.
//TODO: Should it be thread-safe?
type Txn struct {
	//All queries performed by this transaction.
	//Allows for safe storage in Queries.
	mutex   sync.Mutex
	Queries []Query
	counter uint32
	//The actual dgraph transaction.
	txn *dgo.Txn
	//List of schema passed by from the database.
	sch schemaList
}

func (t *Txn) Commit(ctx context.Context) error {
	return t.txn.Commit(ctx)
}

func (t *Txn) Discard(ctx context.Context) error {
	return t.txn.Discard(ctx)
}

func (t *Txn) SetSchema(sch schemaList) {
	t.sch = sch
}

//Perform a single mutation.
func (t *Txn) mutate(ctx context.Context, q Query) error {
	//Add a single mutation to the query list.
	byt, _, err := q.Process(t.sch)
	if err != nil {
		return err
	}
	var m api.Mutation
	if q.Type() == QueryDelete {
		//TODO: fix this
		m.DeleteJson = byt
	} else if q.Type() == QuerySet {
		m.SetJson = byt
	}
	_, err = t.txn.Mutate(ctx, &m)
	return err
}

//Upsert follows the new 1.1 api and performs an upsert.
//TODO: I really don't know what this does so work on it later.
func (t *Txn) Upsert(ctx context.Context, q Query, m []*api.Mutation, obj ...interface{}) error {
	if t.txn == nil {
		return Error(errTransaction)
	}
	b, ma, err := q.Process(t.sch)
	if err != nil {
		return Error(err)
	}
	var req = api.Request{
		//TODO: Dont use create(b) as that performs unnecessary allocations. We do not perform any changes to b which should not cause any issues.
		Query:     bytesToStringUnsafe(b),
		Vars:      ma,
		Mutations: m,
	}
	resp, err := t.txn.Do(ctx, &req)
	if err != nil {
		return Error(err)
	}
	err = HandleResponse(resp.Json, obj)
	return Error(err)
}

func (t *Txn) query(ctx context.Context, q Query, objs []interface{}) error {
	str, m, err := q.Process(t.sch)
	if err != nil {
		return err
	}
	fmt.Println(string(str))
	resp, err := t.txn.QueryWithVars(ctx, bytesToStringUnsafe(str), m)
	if err != nil {
		return Error(err)
	}
	err = HandleResponse(resp.Json, objs)
	if err != nil {
		return Error(err)
	}
	return nil
}

//RunQuery executes the GraphQL+- query.
//If q is a mutation query the mutation objects are supplied in q and not in objs.
func (t *Txn) RunQuery(ctx context.Context, q Query, objs ...interface{}) error {
	//if t.txn == nil {
	//	return Error(errTransaction)
	//}
	//Allow thread-safe appending of queries as might run queries.
	//TODO: Right now this is only for storing the queries. Running queries in parallel that rely on each other is very much a race condition.
	t.mutex.Lock()
	t.Queries = append(t.Queries, q)
	t.mutex.Unlock()
	switch q.Type() {
	case QueryRegular:
		return Error(t.query(ctx, q, objs))
	case QueryDelete, QuerySet:
		return Error(t.mutate(ctx, q))
	}
	return Error(errInvalidType)
}
