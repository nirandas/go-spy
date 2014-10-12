package spy

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

func Match(v []interface{}, m []Matcher) bool {
	if len(v) != len(m) {
		return false
	}
	if len(v) == 0 {
		return true
	}

	for i := 0; i < len(m); i++ {
		if !m[i].Match(v[i]) {
			return false
		}
	}
	return true
}

type Expectation struct {
	funcName  string
	arguments []Matcher
	ret       []interface{}
	calls     []*Call
}

func NewExpectation(funcName string, m ...Matcher) *Expectation {
	return &Expectation{
		funcName:  funcName,
		arguments: m,
	}
}

func (e *Expectation) Return(values ...interface{}) *Expectation {
	e.ret = values
	return e
}

type Call struct {
	name        string
	arguments   []interface{}
	expectation *Expectation
}

func NewCall(e *Expectation, args ...interface{}) *Call {
	return &Call{
		name:        e.funcName,
		arguments:   args,
		expectation: e,
	}
}

func (c *Call) String(i int) string {
	return c.expectation.ret[i].(string)
}

func (c *Call) Int(i int) int {
	return c.expectation.ret[i].(int)
}

func (c *Call) Bool(i int) bool {
	return c.expectation.ret[i].(bool)
}

type Spy struct {
	expectations []*Expectation
	calls        []*Call
}

func (spy *Spy) findMatchingExpectation(funcName string, args ...interface{}) *Expectation {
	for _, e := range spy.expectations {
		if e.funcName == funcName && Match(args, e.arguments) {
			return e
		}
	}
	return nil
}

func (spy *Spy) When(funcName string, matchers ...Matcher) *Expectation {
	e := NewExpectation(funcName, matchers...)
	spy.expectations = append(spy.expectations, e)
	return e
}

func (spy *Spy) Called(args ...interface{}) *Call {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("Couldn't get the caller information")
	}
	functionPath := runtime.FuncForPC(pc).Name()
	parts := strings.Split(functionPath, ".")
	functionName := parts[len(parts)-1]

	e := spy.findMatchingExpectation(functionName, args...)
	if e == nil {
		panic(fmt.Sprintf("Unexpected call to %s with arguments %v", functionName, args))
	}
	c := NewCall(e, args...)
	e.calls = append(e.calls, c)
	spy.calls = append(spy.calls, c)
	return c
}

func (spy *Spy) Verify(t *testing.T) {
	fail := 0
	for _, e := range spy.expectations {
		if len(e.calls) == 0 {
			t.Log(fmt.Sprintf("Expected %s %v to be called", e.funcName, e.arguments))
			fail++
		}
	}

	if fail > 0 {
		t.Fail()
	}
}
