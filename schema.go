package mulbase

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/valyala/fastjson"
)

type SchemaField struct {
	Type    VarType `json:"type,omitempty"`
	List    bool    `json:"list,omitempty"`
	Lang    bool    `json:"lang,omitempty"`
	Reverse bool    `json:"reverse,omitempty"`
	Uid     bool
	Skip    bool
}

var basic = &SchemaField{Skip: true}
var schema = make(map[string]SchemaField)
var wg = sync.WaitGroup{}

//This function runs really early, so check if config has been init.
func initSchema() map[string]SchemaField {
	//TODO: fix this

	if len(schema) > 0 {
		wg.Wait()
		return schema
	}
	wg.Add(1)
	//Set port - 1000
	var res *http.Response
	var err error
	if dotls {
		cer, err := tls.LoadX509KeyPair(tlsPaths[1], tlsPaths[2])
		if err != nil {
			panic(err)
		}
		cc := &tls.Config{Certificates: []tls.Certificate{cer}, InsecureSkipVerify: true}
		tc := &http.Transport{TLSClientConfig: cc}
		var ht = &http.Client{Transport: tc}
		res, err = ht.Post("https://"+ip+":"+strconv.Itoa(port-1000)+"/query", "application/graphql+-", strings.NewReader("schema{}"))
	} else {
		res, err = http.Post("http://"+ip+":"+strconv.Itoa(port-1000)+"/query", "application/graphql+-", strings.NewReader("schema{}"))
	}
	if err != nil {
		panic(err)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	p := fastjson.Parser{}
	n, err := p.ParseBytes(b)
	if err != nil {
		panic(err)
	}
	arr, err := n.GetObject("data").Get("schema").Array()
	if err != nil {
		panic(err)
	}
	for _, v := range arr {
		name := strings.Trim(v.Get("predicate").String(), "\"")
		if name == "_predicate_" {
			continue
		}
		if strings.Contains(name, "dgraph") {
			continue
		}
		if strings.Contains(name, "@") {
			continue
		}
		var b []byte
		b = v.MarshalTo(b)
		var s SchemaField
		DeserializeByte(b, &s)
		s.Uid = s.Type == "uid"
		schema[name] = s
	}
	wg.Done()
	return schema
}

func getSchemaField(name string) *SchemaField {
	str := strings.Split(name, "|")
	if len(str) > 1 {
		return basic
	}
	if name == "uid" {
		return &SchemaField{}
	} else {
		s := schema
		val, ok := s[name]
		if ok {
			return &val
		} else {
			return nil
		}
	}
}
