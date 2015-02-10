package parc

type Grammar struct {
	slots map[string]*GrammarSlot
	start string
}

type MatchFunc func([]byte) (int, bool)

type GrammarSlot struct {
	Type      slotType
	Alts      []*GrammarSlot // for alternation
	Symbol    string         // for non-terminal
	Slot      *GrammarSlot   // for non-terminal
	Continue  *GrammarSlot
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
		slots: make(map[string]*GrammarSlot),
	}
}

func (g *Grammar) Rule(symbol string, slot *GrammarSlot) {
	g.slots[symbol] = slot
}
