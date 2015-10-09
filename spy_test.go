package spy_test

import (
	"errors"
	. "github.com/nirandas/go-spy"

	//	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func Test_new_expectation_creates_new(t *testing.T) {
	RegisterTestingT(t)
	e := NewExpectation("Something")
	Expect(e).ShouldNot(BeNil())
	Expect(e.HasReturns()).To(BeFalse())
}

func Test_expectation_can_set_returns(t *testing.T) {
	RegisterTestingT(t)
	e := NewExpectation("Something")
	e.Return(1, 2)
	Expect(e.HasReturns()).To(BeTrue())
	Expect(e.CountReturns()).To(Equal(2))
}

func Test_call_return_values(t *testing.T) {
	RegisterTestingT(t)
	var c *Call
	e := NewExpectation("Something")
	e.Return("ok", int(1), true, errors.New("dummy"), nil)
	c = NewCall(e)
	Expect(c).ShouldNot(BeNil())
	//can access return as string
	Expect(c.String(0)).To(Equal("ok"))
	//can access return as int
	Expect(c.Int(1)).To(Equal(1))
	//can access return as bool
	Expect(c.Bool(2)).To(Equal(true))
	//can access return as error
	Expect(c.Error(3)).To(HaveOccurred())
	//can access return as nil error
	Expect(c.Error(4)).To(BeNil())
}

func Test_spy_verification_exits_if_calls_missing_(t *testing.T) {
	RegisterTestingT(t)
	s := TestImplementation{}
	//setup expectation
	s.When("DoSomething", String("Hi")).Return("Bye")

	//mock loger for verifying spy.Verify()
	//calls Errorf(...) and Fail() on missing call
	logger := &TB{}
	logger.When("Errorf", Anything(), Anything(), Anything())
	logger.When("Fail")
	defer logger.Verify(t)
	//call s.Verify and pass mock logger
	s.Verify(logger)
}

func Test_it_identifies_proper_expectation(t *testing.T) {
	RegisterTestingT(t)
	s := TestImplementation{}
	//setup expectation
	s.When("DoSomething", String("Hi")).Return("Bye")

	c := s.DoSomething("Hi")
	Expect(c.String(0)).To(Equal("Bye"))
}

func Test_panics_on_unexpected_call(t *testing.T) {
	RegisterTestingT(t)
	s := TestImplementation{}
	//setup expectation
	s.When("DoSomething", String("Hi")).Return("Bye")
	Expect(func() {
		s.DoSomething("Hi", "Hello")
	}).Should(Panic())
}

type TestImplementation struct {
	Spy
}

func (s *TestImplementation) DoSomething(args ...interface{}) *Call {
	return s.Called(args...)
}

type TB struct {
	Spy
}

func (s *TB) Error(args ...interface{}) {
	s.Called(args...)
}
func (s *TB) Errorf(format string, args ...interface{}) {
	s.Called(append([]interface{}{format}, args...)...)
}
func (s *TB) Fail() {
	s.Called()
}
func (s *TB) FailNow() {
	s.Called()
}
func (s *TB) Failed() bool {
	c := s.Called()
	return c.Bool(0)
}

func (s *TB) Fatal(args ...interface{}) {
	s.Called(args...)
}
func (s *TB) Fatalf(format string, args ...interface{}) {
	s.Called(append([]interface{}{format}, args...))
}
func (s *TB) Log(args ...interface{}) {
	s.Called(args...)
}
func (s *TB) Logf(format string, args ...interface{}) {
	s.Called(append([]interface{}{format}, args...))
}
func (s *TB) Skip(args ...interface{}) {
	s.Called(args...)
}
func (s *TB) SkipNow() {
	s.Called()
}
func (s *TB) Skipf(format string, args ...interface{}) {
	s.Called(append([]interface{}{format}, args...))
}
func (s *TB) Skipped() bool {
	c := s.Called()
	return c.Bool(0)
}
