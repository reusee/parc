package parc

type Grammar struct {
	slots map[string]*grammarSlot
	start string
}

type MatchFunc func([]byte) (int, bool)

type grammarSlot struct {
	Type      slotType
	Alts      []*grammarSlot // for alternation
	Symbol    string         // for non-terminal
	Slot      *grammarSlot   // for non-terminal
	Continue  *grammarSlot
	MatchFunc MatchFunc // for terminal
	Name      string    // for debug
}

type slotType int

const (
	slotAlt slotType = iota
	slotNonTerminal
	slotTerminal
	slotReturn
	slotFinish
)

func NewGrammar(startSymbol string) *Grammar {
	return &Grammar{
		start: startSymbol,
		slots: make(map[string]*grammarSlot),
	}
}

func (g *Grammar) Rule(symbol string, slot *grammarSlot) {
	g.slots[symbol] = slot
}
