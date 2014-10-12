package spy_test

import (
	. "github.com/nirandas/go-spy"
	"testing"
)

func TestArgumentsMatching(t *testing.T) {
	a := []interface{}{"Foo", true}
	m := []Matcher{String("Foo"), Bool(true)}

	if !Match(a, m) {
		t.Fatal("Matching failed")
	}

	a[1] = false
	if Match(a, m) {
		t.Fatal("Matching failed")
	}
}

func TestSpying(t *testing.T) {
	s := TestObj{}
	defer s.Spy.Verify(t)
	s.When("DoSomething").Return("Hello bar")
	s.When("DoSomething", String("Hi")).Return("Bye")
	str := s.DoSomething()
	if str != "Hello bar" {
		t.Fatal("Expected Hello bar")
	}
	str = s.DoSomething("Hi")
	if str != "Bye" {
		t.Fatal("Expected 'Bye'")
	}
}

func TestUnexpectedCallPanics(t *testing.T) {
	s := TestObj{}
	defer func() {
		recover()
	}()

	_ = s.DoSomething()
	t.Fatal("Expected panic")
}

type TestObj struct {
	Spy
}

func (s *TestObj) DoSomething(v ...interface{}) string {
	c := s.Spy.Called(v...)
	return c.String(0)
}
