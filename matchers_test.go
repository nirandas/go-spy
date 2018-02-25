package spy

import (
	. "github.com/onsi/gomega"
	"testing"
)

var trueMatcher, falseMatcher, strMatcher, iMatcher, i64Matcher, ui64Matcher, anyMatcher, nullMatcher, notNullMatcher, tMatcher, custMatcher Matcher

func init() {
	trueMatcher = Bool(true)
	falseMatcher = Bool(false)
	strMatcher = String("Foo")
	iMatcher = Int(1)
	i64Matcher = Int64(1)
	ui64Matcher = Uint64(1)
	anyMatcher = Anything()
	nullMatcher = Nil()
	notNullMatcher = NotNil()
	tMatcher = Type("*int")
	custMatcher = Custom(func(a interface{}) bool {
		v, ok := a.(string)
		return ok && len(v) > 4
	})
}

func TestMatchesFoo(t *testing.T) {
	RegisterTestingT(t)
	Expect(strMatcher.Match("Foo")).To(BeTrue(), "failed to match 'Foo' with 'Foo'")
}

func Test_does_not_panic_when_matched_against_non_string(t *testing.T) {
	RegisterTestingT(t)
	Expect(func() {
		strMatcher.Match(2)
	}).ShouldNot(Panic(), "String matcher should not panic")
}

func Test_does_not_match_bar(t *testing.T) {
	RegisterTestingT(t)
	Expect(strMatcher.Match("bar")).To(BeFalse())
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
		Expect(iMatcher.Match(uint(1))).To(BeFalse())
	}).ShouldNot(Panic())
}

func Test_int_matches_1(t *testing.T) {
	RegisterTestingT(t)
	Expect(iMatcher.Match(1)).To(BeTrue())
}

func Test_does_not_match_2(t *testing.T) {
	RegisterTestingT(t)
	Expect(iMatcher.Match(2)).To(BeFalse())
}

func Test_does_not_match_uint_or_int32_or_int64(t *testing.T) {
	RegisterTestingT(t)
	Expect(iMatcher.Match(uint(1))).To(BeFalse())
	Expect(iMatcher.Match(int64(1))).To(BeFalse())
	Expect(iMatcher.Match(int32(1))).To(BeFalse())
}

func Test_does_not_panic_when_matched_against_non_int64(t *testing.T) {
	RegisterTestingT(t)
	Expect(func() {
		Expect(i64Matcher.Match(uint(1))).To(BeFalse())
	}).ShouldNot(Panic())
}

func Test_matches1(t *testing.T) {
	RegisterTestingT(t)
	Expect(i64Matcher.Match(int64(1))).To(BeTrue())
}

func Test_does_not_matches2(t *testing.T) {
	RegisterTestingT(t)
	Expect(i64Matcher.Match(int64(2))).To(BeFalse())
}

func Test_does_not_matches_uint_int_int(t *testing.T) {
	RegisterTestingT(t)
	Expect(i64Matcher.Match(uint(1))).To(BeFalse())
	Expect(i64Matcher.Match(int(1))).To(BeFalse())
	Expect(i64Matcher.Match(int32(1))).To(BeFalse())
}

func Test_does_not_panic_when_matchedagainst_no_uint64(t *testing.T) {
	RegisterTestingT(t)
	Expect(func() {
		Expect(ui64Matcher.Match(int(1))).To(BeFalse())
	}).ShouldNot(Panic())
}

func Test_matches_uint1(t *testing.T) {
	RegisterTestingT(t)
	Expect(ui64Matcher.Match(uint64(1))).To(BeTrue())
}

func Test_does_not_matches_uint2(t *testing.T) {
	RegisterTestingT(t)
	Expect(ui64Matcher.Match(uint64(2))).To(BeFalse())
}

func Test_does_not_match_uint_int32_int(t *testing.T) {
	RegisterTestingT(t)
	Expect(ui64Matcher.Match(uint(1))).To(BeFalse())
	Expect(ui64Matcher.Match(int(1))).To(BeFalse())
	Expect(ui64Matcher.Match(int32(1))).To(BeFalse())
}

func Test_matches_anything(t *testing.T) {
	RegisterTestingT(t)
	Expect(anyMatcher.Match(nil)).To(BeTrue())
	Expect(anyMatcher.Match("1")).To(BeTrue())
	Expect(anyMatcher.Match(1)).To(BeTrue())
	Expect(anyMatcher.Match(String("test"))).To(BeTrue())
}

func Test_matches_nil(t *testing.T) {
	RegisterTestingT(t)
	Expect(nullMatcher.Match(nil)).To(BeTrue())
}

func Test_does_not_matches_non_nil(t *testing.T) {
	RegisterTestingT(t)
	Expect(nullMatcher.Match("value")).To(BeFalse())
	Expect(nullMatcher.Match(10)).To(BeFalse())
}

func Test_match_not_nil(t *testing.T) {
	RegisterTestingT(t)
	Expect(notNullMatcher.Match("value")).To(BeTrue())
	Expect(notNullMatcher.Match(10)).To(BeTrue())
}

func Test_does_not_match_nnil(t *testing.T) {
	RegisterTestingT(t)
	Expect(notNullMatcher.Match(nil)).To(BeFalse())
}

func Test_match_pointer_int(t *testing.T) {
	RegisterTestingT(t)
	var i int
	Expect(tMatcher.Match(&i)).To(BeTrue())
}

func Test_does_not_match_no_pointer_int(t *testing.T) {
	RegisterTestingT(t)
	Expect(tMatcher.Match(1)).To(BeFalse())
	Expect(tMatcher.Match("1")).To(BeFalse())
	Expect(tMatcher.Match(nil)).To(BeFalse())
}

func Test_match_foobar(t *testing.T) {
	RegisterTestingT(t)
	Expect(custMatcher.Match("foobar")).To(BeTrue())
}

func Test_does_not_match_non_string_or_string_lesthan_4_len(t *testing.T) {
	RegisterTestingT(t)
	Expect(custMatcher.Match(1)).To(BeFalse())
	Expect(custMatcher.Match("1")).To(BeFalse())
	Expect(custMatcher.Match(nil)).To(BeFalse())
}
