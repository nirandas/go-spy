package spy_test

import (
	"errors"
	. "github.com/nirandas/go-spy"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Expectation", func() {
	var e *Expectation
	BeforeEach(func() {
		e = NewExpectation("Something")
		Expect(e).ShouldNot(BeNil())
	})
	It("sets returns", func() {
		Expect(e.HasReturns()).To(BeFalse())
		e.Return(1, 2)
		Expect(e.HasReturns()).To(BeTrue())
	})
})

var _ = Describe("Call", func() {
	var c *Call
	BeforeEach(func() {
		e := NewExpectation("Something")
		e.Return("ok", int(1), true, errors.New("dummy"), nil)
		c = NewCall(e)
		Expect(c).ShouldNot(BeNil())
	})

	It("Has 5 returns", func() {
		Expect(c.CountReturns()).To(Equal(5))
	})

	It("can access return as string", func() {
		Expect(c.String(0)).To(Equal("ok"))
	})

	It("can access return as int", func() {
		Expect(c.Int(1)).To(Equal(1))
	})

	It("can access return as bool", func() {
		Expect(c.Bool(2)).To(Equal(true))
	})

	It("can access return as error", func() {
		Expect(c.Error(3)).To(HaveOccurred())
	})

	It("can access return as nil error", func() {
		Expect(c.Error(4)).To(BeNil())
	})

})

var _ = Describe("Spy", func() {
	var s TestImplementation
	BeforeEach(func() {
		s = TestImplementation{}
	})

	Describe("with expectations", func() {
		BeforeEach(func() {
			e := s.When("DoSomething", String("Hi"))
			Expect(e).ShouldNot(BeNil())
			e.Return("Bye")
		})

		It("calls Errorf(...) and Fail() on missing call", func() {
			t := &TB{}
			t.When("Errorf", Anything(), Anything(), Anything())
			t.When("Fail")
			defer t.Verify(GinkgoT())

			s.Verify(t)
		})

		It("identifies proper expectation", func() {
			c := s.DoSomething("Hi")
			Expect(c.String(0)).To(Equal("Bye"))
		})

		It("panics on unexpected call", func() {
			Expect(func() {
				s.DoSomething("Hi", "Hello")
			}).Should(Panic())
		})

	})

})

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
