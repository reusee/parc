package parc

func Alt(alts ...*GrammarSlot) *GrammarSlot {
	return &GrammarSlot{
		Type: slotAlt,
		Alts: alts,
	}
}

func Seq(slots ...*GrammarSlot) (ret *GrammarSlot) {
	if len(slots) == 0 {
		return &GrammarSlot{
			Type: slotReturn,
		}
	}
	ret = slots[0]
	ret.Continue = Seq(slots[1:]...)
	return
}

func N(symbol string) *GrammarSlot {
	return &GrammarSlot{
		Type:   slotNonTerminal,
		Symbol: symbol,
	}
}

func T(matchFunc MatchFunc) *GrammarSlot {
	return &GrammarSlot{
		Type:      slotTerminal,
		MatchFunc: matchFunc,
		Continue: &GrammarSlot{
			Type: slotReturn,
		},
	}
}

func Named(name string, slot *GrammarSlot) *GrammarSlot {
	slot.Name = name
	return slot
}
