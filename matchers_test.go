package spy_test

import (
	. "github.com/nirandas/go-spy"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Matchers", func() {

	Describe("String matcher", func() {
		s := String("Foo")

		It("Does not panic when matched against non-string", func() {
			Expect(func() {
				s.Match(2)
			}).ShouldNot(Panic())
		})

		It("Matches 'Foo'", func() {
			Expect(s.Match("Foo")).To(BeTrue())
		})

		It("Does not match 'bar'", func() {
			Expect(s.Match("bar")).To(BeFalse())
		})

	})

	Describe("Boolng matcher", func() {
		t := Bool(true)
		f := Bool(false)

		It("Does not panic when matched against non-bool", func() {
			Expect(func() {
				t.Match("a")
				f.Match("a")
			}).ShouldNot(Panic())
		})

		It("matches", func() {
			Expect(t.Match(true)).To(BeTrue())
			Expect(f.Match(false))
		})

		It("Does not match", func() {
			Expect(t.Match(false)).To(BeFalse())
			Expect(f.Match(true)).To(BeFalse())
		})

	})

	Describe("int matcher", func() {
		i := Int(1)

		It("Does not panic when matched against non-int", func() {
			Expect(func() {
				Expect(i.Match(uint(1))).To(BeFalse())
			}).ShouldNot(Panic())
		})

		It("Matches 1", func() {
			Expect(i.Match(1)).To(BeTrue())
		})

		It("Does not match 2", func() {
			Expect(i.Match(2)).To(BeFalse())
		})

		It("Does not match uint or int32 or int64", func() {
			Expect(i.Match(uint(1))).To(BeFalse())
			Expect(i.Match(int64(1))).To(BeFalse())
			Expect(i.Match(int32(1))).To(BeFalse())
		})

	})

	Describe("int64 matcher", func() {
		i := Int64(1)

		It("Does not panic when matched against non-int", func() {
			Expect(func() {
				Expect(i.Match(uint(1))).To(BeFalse())
			}).ShouldNot(Panic())
		})

		It("Matches 1", func() {
			Expect(i.Match(int64(1))).To(BeTrue())
		})

		It("Does not match 2", func() {
			Expect(i.Match(int64(2))).To(BeFalse())
		})

		It("Does not match uint or int32 or int", func() {
			Expect(i.Match(uint(1))).To(BeFalse())
			Expect(i.Match(int(1))).To(BeFalse())
			Expect(i.Match(int32(1))).To(BeFalse())
		})

	})

	Describe("uint64 matcher", func() {
		i := Uint64(1)

		It("Does not panic when matched against non-uint64", func() {
			Expect(func() {
				Expect(i.Match(int(1))).To(BeFalse())
			}).ShouldNot(Panic())
		})

		It("Matches 1", func() {
			Expect(i.Match(uint64(1))).To(BeTrue())
		})

		It("Does not match 2", func() {
			Expect(i.Match(uint64(2))).To(BeFalse())
		})

		It("Does not match uint or int32 or int", func() {
			Expect(i.Match(uint(1))).To(BeFalse())
			Expect(i.Match(int(1))).To(BeFalse())
			Expect(i.Match(int32(1))).To(BeFalse())
		})

	})

	Describe("Anything matcher", func() {
		m := Anything()
		It("Matches anything", func() {
			Expect(m.Match(nil)).To(BeTrue())
			Expect(m.Match("1")).To(BeTrue())
			Expect(m.Match(1)).To(BeTrue())
			Expect(m.Match(String("test"))).To(BeTrue())
		})
	})

	Describe("Nil matcher", func() {
		m := Nil()
		It("Matches nil", func() {
			Expect(m.Match(nil)).To(BeTrue())
		})
		It("Does not matches not nil", func() {
			Expect(m.Match("value")).To(BeFalse())
			Expect(m.Match(10)).To(BeFalse())
		})
	})

	Describe("NotNil matcher", func() {
		m := NotNil()
		It("Match not nil", func() {
			Expect(m.Match("value")).To(BeTrue())
			Expect(m.Match(10)).To(BeTrue())
		})
		It("Does not Matches nil", func() {
			Expect(m.Match(nil)).To(BeFalse())
		})
	})

	Describe("Type matcher", func() {
		m := Type("*int")
		It("Match *int", func() {
var i int
			Expect(m.Match(&i)).To(BeTrue())
		})
		It("Does not Matche non-*int", func() {
			Expect(m.Match(1)).To(BeFalse())
			Expect(m.Match("1")).To(BeFalse())
			Expect(m.Match(nil)).To(BeFalse())
		})
	})

	Describe("Custom matcher", func() {
		m := Custom(func(a interface{})bool{
v,ok:=a.(string)
return ok && len(v)>4
})
		It("Match 'foobar'", func() {
			Expect(m.Match("foobar")).To(BeTrue())
		})
		It("Does not Matche non-string or strings less than 4", func() {
			Expect(m.Match(1)).To(BeFalse())
			Expect(m.Match("1")).To(BeFalse())
			Expect(m.Match(nil)).To(BeFalse())
		})
	})

})
