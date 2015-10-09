package spy_test

import (
	. "github.com/nirandas/go-spy"

	//	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var trueMatcher, falseMatcher, stringMatcher, intMatcher, int64Matcher, uint64Matcher, anythingMatcher, nilMatcher, notNilMatcher, typeMatcher, customMatcher Matcher

func init() {
	trueMatcher = Bool(true)
	falseMatcher = Bool(false)
	stringMatcher = String("Foo")
	intMatcher = Int(1)
	int64Matcher = Int64(1)
	uint64Matcher = Uint64(1)
	anythingMatcher = Anything()
	nilMatcher = Nil()
	notNilMatcher = NotNil()
	typeMatcher = Type("*int")
	customMatcher = Custom(func(a interface{}) bool {
		v, ok := a.(string)
		return ok && len(v) > 4
	})
}

func TestMatchesFoo(t *testing.T) {
	RegisterTestingT(t)
	Expect(stringMatcher.Match("Foo")).To(BeTrue(), "failed to match 'Foo' with 'Foo'")
}

func Test_does_not_panic_when_matched_against_non_string(t *testing.T) {
	RegisterTestingT(t)
	Expect(func() {
		stringMatcher.Match(2)
	}).ShouldNot(Panic(), "String matcher should not panic")
}

func Test_does_not_match_bar(t *testing.T) {
	RegisterTestingT(t)
	Expect(stringMatcher.Match("bar")).To(BeFalse())
}

func Test_does_not_panic_when_matched_against_non_bool(t *testing.T) {
	Expect(func() {
		trueMatcher.Match("a")
		falseMatcher.Match("a")
	}).ShouldNot(Panic())
}

func Test_bool_matches(t *testing.T) {
	RegisterTestingT(t)
	Expect(trueMatcher.Match(true)).To(BeTrue())
	Expect(falseMatcher.Match(false))
}

func Test_bool_does__not_match(t *testing.T) {
	RegisterTestingT(t)
	Expect(trueMatcher.Match(false)).To(BeFalse())
	Expect(falseMatcher.Match(true)).To(BeFalse())
}

func Test_does_not_panic_when_matched_against_non_int(t *testing.T) {
	RegisterTestingT(t)
	Expect(func() {
		Expect(intMatcher.Match(uint(1))).To(BeFalse())
	}).ShouldNot(Panic())
}

func Test_int_matches_1(t *testing.T) {
	RegisterTestingT(t)
	Expect(intMatcher.Match(1)).To(BeTrue())
}

func Test_does_not_match_2(t *testing.T) {
	RegisterTestingT(t)
	Expect(intMatcher.Match(2)).To(BeFalse())
}

func Test_does_not_match_uint_or_int32_or_int64(t *testing.T) {
	RegisterTestingT(t)
	Expect(intMatcher.Match(uint(1))).To(BeFalse())
	Expect(intMatcher.Match(int64(1))).To(BeFalse())
	Expect(intMatcher.Match(int32(1))).To(BeFalse())
}

func Test_does_not_panic_when_matched_against_non_int64(t *testing.T) {
	RegisterTestingT(t)
	Expect(func() {
		Expect(int64Matcher.Match(uint(1))).To(BeFalse())
	}).ShouldNot(Panic())
}

func Test_matches1(t *testing.T) {
	RegisterTestingT(t)
	Expect(int64Matcher.Match(int64(1))).To(BeTrue())
}

func Test_does_not_matches2(t *testing.T) {
	RegisterTestingT(t)
	Expect(int64Matcher.Match(int64(2))).To(BeFalse())
}

func Test_does_not_matches_uint_int_int(t *testing.T) {
	RegisterTestingT(t)
	Expect(int64Matcher.Match(uint(1))).To(BeFalse())
	Expect(int64Matcher.Match(int(1))).To(BeFalse())
	Expect(int64Matcher.Match(int32(1))).To(BeFalse())
}

func Test_does_not_panic_when_matchedagainst_no_uint64(t *testing.T) {
	RegisterTestingT(t)
	Expect(func() {
		Expect(uint64Matcher.Match(int(1))).To(BeFalse())
	}).ShouldNot(Panic())
}

func Test_matches_uint1(t *testing.T) {
	RegisterTestingT(t)
	Expect(uint64Matcher.Match(uint64(1))).To(BeTrue())
}

func Test_does_not_matches_uint2(t *testing.T) {
	RegisterTestingT(t)
	Expect(uint64Matcher.Match(uint64(2))).To(BeFalse())
}

func Test_does_not_match_uint_int32_int(t *testing.T) {
	RegisterTestingT(t)
	Expect(uint64Matcher.Match(uint(1))).To(BeFalse())
	Expect(uint64Matcher.Match(int(1))).To(BeFalse())
	Expect(uint64Matcher.Match(int32(1))).To(BeFalse())
}

func Test_matches_anything(t *testing.T) {
	RegisterTestingT(t)
	Expect(anythingMatcher.Match(nil)).To(BeTrue())
	Expect(anythingMatcher.Match("1")).To(BeTrue())
	Expect(anythingMatcher.Match(1)).To(BeTrue())
	Expect(anythingMatcher.Match(String("test"))).To(BeTrue())
}

func Test_matches_nil(t *testing.T) {
	RegisterTestingT(t)
	Expect(nilMatcher.Match(nil)).To(BeTrue())
}

func Test_does_not_matches_non_nil(t *testing.T) {
	RegisterTestingT(t)
	Expect(nilMatcher.Match("value")).To(BeFalse())
	Expect(nilMatcher.Match(10)).To(BeFalse())
}

func Test_match_not_nil(t *testing.T) {
	RegisterTestingT(t)
	Expect(notNilMatcher.Match("value")).To(BeTrue())
	Expect(notNilMatcher.Match(10)).To(BeTrue())
}

func Test_does_not_match_nnil(t *testing.T) {
	RegisterTestingT(t)
	Expect(notNilMatcher.Match(nil)).To(BeFalse())
}

func Test_match_pointer_int(t *testing.T) {
	RegisterTestingT(t)
	var i int
	Expect(typeMatcher.Match(&i)).To(BeTrue())
}

func Test_does_not_match_no_pointer_int(t *testing.T) {
	RegisterTestingT(t)
	Expect(typeMatcher.Match(1)).To(BeFalse())
	Expect(typeMatcher.Match("1")).To(BeFalse())
	Expect(typeMatcher.Match(nil)).To(BeFalse())
}

func Test_match_foobar(t *testing.T) {
	RegisterTestingT(t)
	Expect(customMatcher.Match("foobar")).To(BeTrue())
}

func Test_does_not_match_non_string_or_string_lesthan_4_len(t *testing.T) {
	RegisterTestingT(t)
	Expect(customMatcher.Match(1)).To(BeFalse())
	Expect(customMatcher.Match("1")).To(BeFalse())
	Expect(customMatcher.Match(nil)).To(BeFalse())
}
