package spy
import(
"reflect"
)

type Matcher interface {
	//Match matches provided value against conditions and returns true or false
	// this method should not panic
	Match(actual interface{}) bool
}

type stringMatcher struct {
	value string
}

func (m stringMatcher) Match(a interface{}) bool {
	v, ok := a.(string)
	return ok && m.value == v
}

//String returns a matcher which match against provided string value
func String(s string) Matcher {
	return stringMatcher{s}
}

type boolMatcher struct {
	value bool
}

func (m boolMatcher) Match(a interface{}) bool {
	v, ok := a.(bool)
	return ok && m.value == v
}

//Bool returns a matcher which can match against provided bool value
func Bool(b bool) Matcher {
	return boolMatcher{b}
}

type anythingMatcher struct{}

func (m anythingMatcher) Match(a interface{}) bool { return true }

//Anything returns a matcher which matches anything
func Anything() Matcher {
	return anythingMatcher{}
}

type nilMatcher struct{}

func (m nilMatcher) Match(a interface{}) bool { return a == nil }

//Nil returns a matcher which matches nil
func Nil() Matcher { return nilMatcher{} }

type notNilMatcher struct{}

func (m notNilMatcher) Match(a interface{}) bool { return a != nil }

//NotNil returns a matcher which matches anything not nil
func NotNil() Matcher { return notNilMatcher{} }

type intMatcher struct {
	value int
}

func (m intMatcher) Match(v interface{}) bool {
	v, ok := v.(int)
	return ok && m.value == v
}

//Int returns a matcher which can match against provided int value
func Int(i int) Matcher { return intMatcher{i} }

type int64Matcher struct {
	value int64
}

func (m int64Matcher) Match(v interface{}) bool {
	v, ok := v.(int64)
	return ok && m.value == v
}

//Int64 returns a matcher which can match against provided int64 value
func Int64(i int64) Matcher { return int64Matcher{i} }

type uint64Matcher struct {
	value uint64
}

func (m uint64Matcher) Match(v interface{}) bool {
	v, ok := v.(uint64)
	return ok && m.value == v
}

//Uint64 returns a matcher which can match against provided uint64 value
func Uint64(i uint64) Matcher { return uint64Matcher{i} }

type typeMatcher struct{
value string
}

//Type returns a Matcher which matches the type against the string argument.
// Example Type("*int")
// Will match a pointer to int
func (m typeMatcher) Match(v interface{})bool{
if v==nil{
return false
}
return reflect.TypeOf(v).String() == m.value
}

func Type(t string)Matcher{
return typeMatcher{t}
}

type customMatcher struct{
value func(v interface{})bool
}

func (m customMatcher) Match(v interface{})bool{
return m.value(v)
}

//Custom wrapps a custom func with Matcher interface
func Custom(f func(v interface{})bool)Matcher{
return customMatcher{f}
}
