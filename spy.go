//spy Provides ability to spy on calls made to an interface
package spy

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type Logger interface {
	Fail()
	Errorf(format string, args ...interface{})
}

//Match matches a slice of v interface{} against a slice of m Matcher interface
//and returns true if each element in the v passes the match against the same element in m
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
	set       map[int]interface{}
}

func NewExpectation(funcName string, m ...Matcher) *Expectation {
	return &Expectation{
		funcName:  funcName,
		arguments: m,
		set:       make(map[int]interface{}),
	}
}

//Return Set the return values for the call
func (e *Expectation) Return(values ...interface{}) *Expectation {
	e.ret = values
	return e
}

func (e *Expectation) Set(index int, v interface{}) *Expectation {
	e.set[index] = v
	return e
}

//HasReturns HasReturns returns true if expectation has return values
func (e *Expectation) HasReturns() bool {
	return len(e.ret) > 0
}

//CountReturn returns number of return values
func (e *Expectation) CountReturns() int {
	return len(e.ret)
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

//CountReturns Returns number of return values
func (c *Call) CountReturns() int {
	return c.expectation.CountReturns()
}

//String Returns the return value at the provided index as a string. Will panic if out of bounds or value can not be converted to string
func (c *Call) String(i int) string {
	return c.expectation.ret[i].(string)
}

func (c *Call) Int(i int) int {
	return c.expectation.ret[i].(int)
}

//Bool Returns the return value at the provided index as a bool . Will panic if out of bounds or value can not be converted to bool
func (c *Call) Bool(i int) bool {
	return c.expectation.ret[i].(bool)
}

//Error Returns the return value at the provided index as a error. Will panic if out of bounds or value can not be converted to error. Handles nill error values gracefully.
func (c *Call) Error(i int) error {
	if c.expectation.ret[i] == nil {
		return nil
	}
	return c.expectation.ret[i].(error)
}

//Get returns the return value at the provided index
func (c *Call) Get(i int) interface{} {
	return c.expectation.ret[i]
}

func (c *Call) GetArg(i int) interface{} {
	return c.arguments[i]
}

func (c *Call) GetReturnValue(index int) interface{}{
	return c.expectation.ret[index]
}

//Spy provides call spying functionalities
//Should be embed in the struct to be mocked
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

//When sets up expectations
//First argument should be the method name and following arguments should be matchers which can be used to match arguments of call
func (spy *Spy) When(funcName string, matchers ...Matcher) *Expectation {
	e := NewExpectation(funcName, matchers...)
	spy.expectations = append(spy.expectations, e)
	return e
}

//Called records the call and returns the Call struct which can be used to access return values
//this is to be called in each of the interface method to be spied
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

	if len(e.set) > 0 {
		for index, set := range e.set {
			if index < 0 || index >= len(args) {
				continue
			}
			value := reflect.ValueOf(args[index])
			if value.Kind() != reflect.Ptr && value.Kind() != reflect.Interface {
				continue
			}
			value = value.Elem()
			if value.CanSet() {
				value.Set(reflect.ValueOf(set))
			}
		}
	}

	return c
}

//Verify verifies all expectations were met
//pass a struct implementing the Logger  interface. The testing.T in golang's testing package satisfies the Logger interface
func (spy *Spy) Verify(t Logger) {
	fail := 0
	for _, e := range spy.expectations {
		if len(e.calls) == 0 {
			t.Errorf("Expected %s %v to be called", e.funcName, e.arguments)
			fail++
		}
	}

	if fail > 0 {
		t.Fail()
	}
}

func (spy *Spy) GetCallsOf(funcName string) []*Call {
	matchingCalls := []*Call{}
	for _,v := range spy.calls {
		if v.name == funcName {
			matchingCalls = append(matchingCalls, v)
		}
	}
	return matchingCalls
}

func (spy *Spy) GetCall(funcName string, index int) *Call{
	return spy.GetCallsOf(funcName)[index]
}

func (spy *Spy) CallCount(funcName string) int {
	return len(spy.GetCallsOf(funcName))
}
