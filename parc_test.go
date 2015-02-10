package parc

import "testing"

func TestMatch(t *testing.T) {
	grammar := &Grammar{
		Start: "S",
		slots: map[string]*grammarSlot{
			"S": &grammarSlot{
				Id:   "S",
				Type: slotAlt,
				Alts: []*grammarSlot{
					&grammarSlot{
						Id:     "S1",
						Type:   slotNonTerminal,
						Symbol: "A",
						Continue: &grammarSlot{
							Id:     "S2",
							Type:   slotNonTerminal,
							Symbol: "S",
							Continue: &grammarSlot{
								Id:        "S3",
								Type:      slotTerminal,
								MatchFunc: ByteEq('d'),
								Continue: &grammarSlot{
									Id:   "S4",
									Type: slotReturn,
								},
							},
						},
					},
					&grammarSlot{
						Id:     "S5",
						Type:   slotNonTerminal,
						Symbol: "B",
						Continue: &grammarSlot{
							Id:     "S6",
							Type:   slotNonTerminal,
							Symbol: "S",
							Continue: &grammarSlot{
								Id:   "S7",
								Type: slotReturn,
							},
						},
					},
					&grammarSlot{
						Id:   "S8",
						Type: slotReturn,
					},
				},
			},
			"A": &grammarSlot{
				Id:        "S9",
				Type:      slotTerminal,
				MatchFunc: ByteIn([]byte{'a', 'c'}),
				Continue: &grammarSlot{
					Id:   "Sa",
					Type: slotReturn,
				},
			},
			"B": &grammarSlot{
				Id:        "Sb",
				Type:      slotTerminal,
				MatchFunc: ByteIn([]byte{'a', 'b'}),
				Continue: &grammarSlot{
					Id:   "Sc",
					Type: slotReturn,
				},
			},
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
