package parc

func Alt(alts ...*grammarSlot) *grammarSlot {
	return &grammarSlot{
		Type: slotAlt,
		Alts: alts,
	}
}

func Seq(slots ...*grammarSlot) (ret *grammarSlot) {
	if len(slots) == 0 {
		return &grammarSlot{
			Type: slotReturn,
		}
	}
	ret = slots[0]
	ret.Continue = Seq(slots[1:]...)
	return
}

func N(symbol string) *grammarSlot {
	return &grammarSlot{
		Type:   slotNonTerminal,
		Symbol: symbol,
	}
}

func T(matchFunc MatchFunc) *grammarSlot {
	return &grammarSlot{
		Type:      slotTerminal,
		MatchFunc: matchFunc,
		Continue: &grammarSlot{
			Type: slotReturn,
		},
	}
}

func Named(name string, slot *grammarSlot) *grammarSlot {
	slot.Name = name
	return slot
}
