package parc

import "testing"

func TestMatch(t *testing.T) {
	grammar := &Grammar{
		Start: "S",
		slots: map[string]*grammarSlot{
			"S": Alt(
				Seq(
					Named("S1", N("A")),
					Named("S2", N("S")),
					Named("S3", T(ByteEq('d'))),
				),
				Seq(
					Named("S5", N("B")),
					Named("S6", N("S")),
				),
				Seq(),
			),
			"A": Named("S9", T(ByteIn([]byte{'a', 'c'}))),
			"B": Named("Sb", T(ByteIn([]byte{'a', 'b'}))),
		},
	}

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
	} {
		if !grammar.Match([]byte(input)) {
			debug = true
			grammar.Match([]byte(input))
			debug = false
			t.Fatalf("match error %s\n", input)
		}
	}
}
