package parc

import "testing"

func TestMatch(t *testing.T) {
	grammar := NewGrammar("S")
	grammar.Rule("S", Named("S", Alt(
		Named("Sa", Seq(
			Named("S1", N("A")),
			Named("S2", N("S")),
			Named("S3", T(ByteEq('d'))),
		)),
		Named("Sb", Seq(
			Named("S4", N("B")),
			Named("S5", N("S")),
		)),
		Named("Sc", Seq()),
	)))
	grammar.Rule("A", Named("A", T(ByteIn([]byte{'a', 'c'}))))
	grammar.Rule("B", Named("B", T(ByteIn([]byte{'a', 'b'}))))

	for _, input := range []string{
		"",
		"a",
		"b",
		"aa",
		"ba",
		"ab",
		"bb",
		"aaad",
		"abad",
		"aabd",
		"abbd",
		"aaa",
		"aba",
		"aab",
		"abb",
		"aaaad",
		"aabad",
		"aaabd",
		"aabbd",
		"baa",
		"bba",
		"bab",
		"bbb",
		"baaad",
		"babad",
		"baabd",
		"babbd",
		"foo",
	} {
		if !grammar.Match([]byte(input)) {
			debug = true
			grammar.Match([]byte(input))
			debug = false
			t.Fatalf("match error %s\n", input)
		}
	}
}
