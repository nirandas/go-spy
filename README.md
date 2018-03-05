go-spy
======

Tiny mocking package for golang

This package provides ability to implement mocks and set and verify expectations. This is an early version, there are hardly any tests or documentation. I hope to improve (add tests, provide documentation, add missing features) and use this package in production soon.

## Example use ##

```go
import (
	"testing"

	. "github.com/nirandas/go-spy"
)

func TestSpying(t *testing.T) {
	s := TestObj{}//mock object
	defer s.Spy.Verify(t)//verifies that expectations were met

	s.When("DoSomething").Return("Hello bar")// set up expectation
	s.When("DoSomething", String("Hi")).Return("Bye")// set up expectation

	str := s.DoSomething()
	if str != "Hello bar" {
		t.Fatal("Expected Hello bar")
	}
	str = s.DoSomething("Hi")
	if str != "Bye" {
		t.Fatal("Expected 'Bye'")
	}
}

//fake implementation
type TestObj struct {
	Spy
}

func (s *TestObj) DoSomething(v ...interface{}) string {
	c := s.Spy.Called(v...)
	return c.String(0)
}
```

## Contributors ##
* [Nirandas Thavorath](https://github.com/nirandas)
* [Eugene Kostrikov](https://github.com/EugeneKostrikov)
