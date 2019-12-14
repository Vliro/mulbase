package humus

import (
	"errors"
	"strconv"
	"strings"
)

/*
var pool sync.Pool

func init() {
	pool.New = func() interface{} {
		return new(GeneratedQuery)
	}
}
*/
/*
	UID represents the primary UID class used in communication with DGraph.
	This is used in code generation.
*/
type UID string

func (u UID) Int() int64 {
	if len(u) < 2 {
		return -1
	}
	val, err := strconv.ParseInt(string(u[2:]), 16, 64)
	if err != nil {
		return -1
	}
	return val
}

func (u UID) IntString() string {
	val, err := strconv.ParseInt(string(u), 16, 64)
	if err != nil {
		return ""
	}
	return strconv.FormatInt(val, 10)
}

func stringFromInt(id int64) string {
	return "0x" + strconv.FormatInt(id, 16)
}

//ParseUid parses a uid from an int64.
func ParseUid(id int64) UID {
	return UID(stringFromInt(id))
}

const (
	// syntax tokens
	tokenLB     = "{" // Left Brace
	tokenRB     = "}" // Right Brace
	tokenLP     = "(" // Left parenthesis
	tokenRP     = ")" // Right parenthesis
	tokenColumn = ":"
	tokenComma  = ","
	tokenSpace  = " "
	tokenFilter = "@filter"
)

type Queries struct {
	q          []*GeneratedQuery
	varCounter func() int
	currentVar int
	vars       map[string]string
}

//Satisfy the Query interface.
func (q *Queries) Process() (string, error) {
	return q.create()
}

func (q *Queries) names() []string {
	c := len(q.q)
	for _, v := range q.q {
		if v.variable.varQuery {
			c--
		}
	}
	ret := make([]string, c)
	count := 0
	for _, v := range q.q {
		if v.variable.varQuery {
			continue
		}
		ret[count] = "q" + strconv.Itoa(v.index)
		count++
	}
	return ret
}

func (q *Queries) NewQuery(f Fields) *GeneratedQuery {
	newq := &GeneratedQuery{
		modifiers: make(map[Predicate]*mapElement),
		fields:    f,
		varMap:    q.vars,
	}
	newq.varFunc = q.varCounter
	q.q = append(q.q, newq)
	return newq
}

//create the byte representation.
func (q *Queries) create() (string, error) {
	var final strings.Builder
	final.Grow(512)
	//The query variable information. Named per default.
	//TODO: Right now it breaks if no queries have vars.
	final.WriteString("query t")
	//The global variable counter. It exists in a closure, it's just easy.
	final.WriteString("(")
	for k, qu := range q.q {
		qu.mapVariables(qu)
		str := qu.variables()
		if str == "" {
			continue
		}
		//TODO: Make it more like a strings.Join to avoid all these error-prone additions.
		if len(q.q) > 1 && k > 0 {
			final.WriteByte(',')
		}
		final.WriteString(str)
	}
	final.WriteByte(')')
	count := 0
	for _, qu := range q.q {
		final.WriteByte('{')
		if !qu.variable.varQuery {
			qu.index = count
			count++
		}
		_, err := qu.create(&final)
		if err != nil {
			return "", err
		}
		final.WriteByte('}')
	}
	return final.String(), nil
}

func (q *Queries) queryVars() map[string]string {
	return q.vars
}

//GeneratedQuery is the root object of queries that are constructed.
//It is constructed using a set of Fields that are either autogenerated or
//manually specified. From its list of modifiers(orderby, pagination etc)
//it automatically and efficiently builds up a query ready for sending to Dgraph.
type GeneratedQuery struct {
	//The root function.
	//Since all queries have a graph function this is an embedded struct.
	//(It is embedded for convenient access as well as unnecessary pointer traversal).
	function
	//Builder for GraphQL variables.
	varBuilder strings.Builder
	//Which directives to apply on this query, such as @cascade.
	directives []Directive
	//The list of fields used in this query.
	fields Fields
	//The overall language for this query.
	language Language
	//List of modifiers, i.e. order, pagination etc.
	modifiers map[Predicate]*mapElement
	//Map for dealing with GraphQL variables. It is inherited in multi-query layout.
	varMap map[string]string
	//function for getting next query value in multi-query. It is also used
	//to define whether it is a single query.
	varFunc func() int
	//varCounter is used for single queries, to keep track of the next GraphQL variable to be assigned.
	varCounter int
	//For multiple queries. Used to keep track of the query name.
	index int
	//top level variable name
	variable struct {
		varName  string
		varQuery bool
	}
	//Whether to allow untagged language.
	strictLanguage bool
}

//NewQuery returns a new singular generation query for use
//in building a single query.
func NewQuery(f Fields) *GeneratedQuery {
	return &GeneratedQuery{
		varMap:    make(map[string]string),
		modifiers: make(map[Predicate]*mapElement),
		fields:    f,
	}
}

//NewQueries returns a QueryList used for building
//multiple queries at once.
func NewQueries() *Queries {
	qu := new(Queries)
	qu.q = make([]*GeneratedQuery, 0, 2)
	qu.currentVar = -1
	qu.varCounter = func() int {
		qu.currentVar++
		return qu.currentVar
	}
	qu.vars = make(map[string]string)
	return qu
}

func (q *GeneratedQuery) Process() (string, error) {
	return q.create(nil)
}

//MutationType defines whether a mutation sets or deletes values.
type MutationType string

const (
	MutateDelete MutationType = "delete"
	MutateSet    MutationType = "set"
)

//Default name for query.
var defaultName = []string{"q0"}

func (q *GeneratedQuery) single() bool {
	return q.varFunc == nil

}
func (q *GeneratedQuery) names() []string {
	if q.variable.varQuery {
		return nil
	}
	if q.single() {
		return defaultName
	}
	return []string{"q" + strconv.Itoa(q.index)}
}

func (q *GeneratedQuery) create(sb *strings.Builder) (string, error) {
	//t := time.Now()
	//sb == nil implies this is a solo query. This means we need to map the GraphQL
	//variables beforehand as it is otherwise calculated in the Queries calculation in Queries.create()
	if sb == nil {
		sb = new(strings.Builder)
		q.mapVariables(q)
		sb.Grow(256)
	}
	if err := q.function.check(q); err != nil {
		return "", err
	}
	//Top level modifiers.
	val, ok := q.modifiers[""]
	//Single query.
	if q.single() && q.variable.varQuery {
		return "", errors.New("singular query with var set is invalid")
	}
	if q.single() {
		vars := q.variables()
		if vars == "" {
			sb.WriteString("query{")
		} else {
			sb.WriteString("query t(")
			sb.WriteString(vars)
			sb.WriteString("){")
		}
	}
	//Write query header.
	if q.variable.varQuery {
		if q.variable.varName != "" {
			sb.WriteString(q.variable.varName)
			sb.WriteString(" as ")
		}
		sb.WriteString("var")
	} else {
		sb.WriteByte('q')
		if q.index != 0 {
			sb.WriteString(strconv.Itoa(q.index))
		} else {
			sb.WriteByte('0')
		}
	}
	sb.WriteString(tokenLP + "func" + tokenColumn + tokenSpace)
	err := q.function.create(q, sb)
	if err != nil {
		return "", err
	}
	if ok {
		//Two passes. Before and after parenthesis. That's just how it be.
		val.m.sort()
		err := val.m.runTopLevel(q, 0, modifierFunction, sb)
		if err != nil {
			return "", err
		}
	} else {
		sb.WriteByte(')')
	}
	for _, v := range q.directives {
		sb.WriteByte('@')
		sb.WriteString(string(v))
	}
	sb.WriteByte('{')
	var parentBuf = make([]byte, 0, 64)
	for _, field := range q.fields.Get() {
		if len(field.Name) > 64 {
			//This code should pretty much never execute as a predicate is rarely this large.
			parentBuf = make([]byte, 2*len(field.Name))
		}
		parentBuf = parentBuf[:len(field.Name)]
		copy(parentBuf, field.Name)
		//parentBuf = append(parentBuf, field.Name...)
		err := field.create(q, parentBuf, sb)
		if err != nil {
			return "", err
		}
	}
	if ok {
		err = val.m.runVariables(q, 0, modifierFunction, sb)
		if err != nil {
			return "", err
		}
	}
	//Add default uid to top level field and close query.
	sb.WriteString(" uid" + tokenRB)
	if q.single() {
		sb.WriteByte('}')
	}
	return sb.String(), nil
}

func (q *GeneratedQuery) queryVars() map[string]string {
	return q.varMap
}

//Directive adds a top level directive.
func (q *GeneratedQuery) Directive(dir Directive) *GeneratedQuery {
	for _, v := range q.directives {
		if v == dir {
			return q
		}
	}
	q.directives = append(q.directives, dir)
	return q
}

/*
At allows you to run modifiers at a path. Modifiers include
pagination, sorting, filters among others.
*/
func (q *GeneratedQuery) At(path Predicate, op Operation) *GeneratedQuery {
	val, ok := q.modifiers[path]
	if !ok {
		val = new(mapElement)
		val.q = q
		q.modifiers[path] = val
	}
	var m = (*modifierCreator)(val)
	if op != nil {
		op(m)
	}
	return q
}

/*
Facets sets @facets for the edge specified by path along with all values as specified by op.
This can be  used to fetch facets, store facets in query variables or something in that manner.
For instance, generating a line in the fashion of @facets(value as friendsSince)
will store the facet value 'friendsSince' into the value variable 'value'.
*/
func (q *GeneratedQuery) Facets(path Predicate, op Operation) *GeneratedQuery {
	val, ok := q.modifiers[path]
	if !ok {
		val = new(mapElement)
		val.q = q
		q.modifiers[path] = val
	}
	var f = (*facetCreator)(val)
	if op != nil {
		op(f)
	}
	f.f.active = true
	return q
}

/*
GroupBy allows you to groupBy at a leaf field. Using op specify a list of variables and operations
to be written as aggregation at this level. onWhich specifies what predicate to actually group on.
*/
func (q *GeneratedQuery) GroupBy(path Predicate, onWhich Predicate, op Operation) *GeneratedQuery {
	val, ok := q.modifiers[path]
	if !ok {
		val = new(mapElement)
		val.q = q
		q.modifiers[path] = val
	}
	var g = (*groupCreator)(val)
	if op != nil {
		op(g)
	}
	g.g.p = onWhich
	return q
}

//Language sets the language for the query to apply to all fields.
//If strict do not allow untagged language.
func (q *GeneratedQuery) Language(l Language, strict bool) *GeneratedQuery {
	q.language = l
	q.strictLanguage = strict
	return q
}

//Returns all the query variables for this query in create form.
func (q *GeneratedQuery) variables() string {
	return q.varBuilder.String()
}

//Var sets q as a var query with the variable name name.
//if name is empty it is just a basic var query.
func (q *GeneratedQuery) Var(name string) *GeneratedQuery {
	q.variable.varQuery = true
	q.variable.varName = name
	return q
}

func (q *GeneratedQuery) registerVariable(typ varType, value string) string {
	if q.varBuilder.Len() != 0 {
		q.varBuilder.WriteByte(',')
	} else {
		q.varBuilder.Grow(32)
	}
	var val int
	if q.varFunc != nil {
		val = q.varFunc()
	} else {
		val = q.varCounter
		q.varCounter++
	}
	key := "$" + strconv.Itoa(val)
	q.varBuilder.WriteString(key)
	q.varBuilder.WriteByte(':')
	q.varBuilder.WriteString(string(typ))
	q.varMap[key] = value
	return key
}

//Static create a static query from the generated version.
//Since this is performed at init, panic if the query
//creation does not work.
func (q *GeneratedQuery) Static() StaticQuery {
	str, err := q.create(nil)
	if err != nil {
		panic(err)
	}
	return StaticQuery{
		Query: str,
		vars:  nil,
	}
}

//Function sets the function type for this function. It is used alongside
//variables. variables are automatically mapped to GraphQL variables as a way
//of avoiding SQL injections.
func (q *GeneratedQuery) Function(ft FunctionType) *GeneratedQuery {
	q.function = function{typ: ft}
	return q
}

//Values simple performs the same as Value but for a variadic number of arguments.
func (q *GeneratedQuery) Values(v ...interface{}) *GeneratedQuery {
	q.function.values(v)
	return q
}
