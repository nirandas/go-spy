package spy

type Matcher interface {
	Match(actual interface{}) bool
}

type stringMatcher struct {
	Value string
}

func (m stringMatcher) Match(a interface{}) bool {
	return m.Value == a.(string)
}

func String(s string) stringMatcher {
	return stringMatcher{s}
}

type boolMatcher struct {
	Value bool
}

func (m boolMatcher) Match(a interface{}) bool {
	return m.Value == a.(bool)
}

func Bool(b bool) boolMatcher {
	return boolMatcher{b}
}

type anythingMatcher struct{}

func (m anythingMatcher) Match(a interface{}) bool { return true }

func Anything() anythingMatcher {
	return anythingMatcher{}
}

type nilMatcher struct{}

func (m nilMatcher) Match(a interface{}) bool { return a == nil }
func Nil() nilMatcher                         { return nilMatcher{} }

type notNilMatcher struct{}

func (m notNilMatcher) Match(a interface{}) bool { return a != nil }
func NotNil() notNilMatcher                      { return notNilMatcher{} }

type intMatcher struct {
	Value int
}

func (m intMatcher) Match(v interface{}) bool {
	return m.Value == v.(int)
}

func Int(i int) intMatcher { return intMatcher{i} }

type int64Matcher struct {
	Value int64
}

func (m int64Matcher) Match(v interface{}) bool {
	return m.Value == v.(int64)
}

func Int64(i int64) int64Matcher { return int64Matcher{i} }

type uint64Matcher struct {
	Value uint64
}

func (m uint64Matcher) Match(v interface{}) bool {
	return m.Value == v.(uint64)
}

func Uint64(i uint64) uint64Matcher { return uint64Matcher{i} }
