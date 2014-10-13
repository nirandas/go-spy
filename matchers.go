package spy

type Matcher interface {
	//Match matches provided value against conditions and returns true or false
	// this method should not panic
	Match(actual interface{}) bool
}

type stringMatcher struct {
	Value string
}

func (m stringMatcher) Match(a interface{}) bool {
	v, ok := a.(string)
	return ok && m.Value == v
}

//String returns a matcher which can match against provided string value
func String(s string) stringMatcher {
	return stringMatcher{s}
}

type boolMatcher struct {
	Value bool
}

func (m boolMatcher) Match(a interface{}) bool {
	v, ok := a.(bool)
	return ok && m.Value == v
}

//Bool returns a matcher which can match against provided bool value
func Bool(b bool) boolMatcher {
	return boolMatcher{b}
}

type anythingMatcher struct{}

func (m anythingMatcher) Match(a interface{}) bool { return true }

//Anything returns a matcher which matches anything
func Anything() anythingMatcher {
	return anythingMatcher{}
}

type nilMatcher struct{}

func (m nilMatcher) Match(a interface{}) bool { return a == nil }

//Nil returns a matcher which matches nil
func Nil() nilMatcher { return nilMatcher{} }

type notNilMatcher struct{}

func (m notNilMatcher) Match(a interface{}) bool { return a != nil }

//NotNil returns a matcher which matches anything not nil
func NotNil() notNilMatcher { return notNilMatcher{} }

type intMatcher struct {
	Value int
}

func (m intMatcher) Match(v interface{}) bool {
	v, ok := v.(int)
	return ok && m.Value == v
}

//Int returns a matcher which can match against provided int value
func Int(i int) intMatcher { return intMatcher{i} }

type int64Matcher struct {
	Value int64
}

func (m int64Matcher) Match(v interface{}) bool {
	v, ok := v.(int64)
	return ok && m.Value == v
}

//Int64 returns a matcher which can match against provided int64 value
func Int64(i int64) int64Matcher { return int64Matcher{i} }

type uint64Matcher struct {
	Value uint64
}

func (m uint64Matcher) Match(v interface{}) bool {
	v, ok := v.(uint64)
	return ok && m.Value == v
}

//Uint64 returns a matcher which can match against provided uint64 value
func Uint64(i uint64) uint64Matcher { return uint64Matcher{i} }
