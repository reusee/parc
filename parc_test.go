package parc

import "testing"

func TestMatch(t *testing.T) {
	grammar := &Grammar{
		Start: "S",
		slots: map[string]*grammarSlot{
			"S": &grammarSlot{
				Name: "S",
				Type: slotAlt,
				Alts: []*grammarSlot{
					&grammarSlot{
						Name:   "S1",
						Type:   slotNonTerminal,
						Symbol: "A",
						Continue: &grammarSlot{
							Name:   "S2",
							Type:   slotNonTerminal,
							Symbol: "S",
							Continue: &grammarSlot{
								Name:      "S3",
								Type:      slotTerminal,
								MatchFunc: ByteEq('d'),
								Continue: &grammarSlot{
									Name: "S4",
									Type: slotReturn,
								},
							},
						},
					},
					&grammarSlot{
						Name:   "S5",
						Type:   slotNonTerminal,
						Symbol: "B",
						Continue: &grammarSlot{
							Name:   "S6",
							Type:   slotNonTerminal,
							Symbol: "S",
							Continue: &grammarSlot{
								Name: "S7",
								Type: slotReturn,
							},
						},
					},
					&grammarSlot{
						Name: "S8",
						Type: slotReturn,
					},
				},
			},
			"A": &grammarSlot{
				Name:      "S9",
				Type:      slotTerminal,
				MatchFunc: ByteIn([]byte{'a', 'c'}),
				Continue: &grammarSlot{
					Name: "Sa",
					Type: slotReturn,
				},
			},
			"B": &grammarSlot{
				Name:      "Sb",
				Type:      slotTerminal,
				MatchFunc: ByteIn([]byte{'a', 'b'}),
				Continue: &grammarSlot{
					Name: "Sc",
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
