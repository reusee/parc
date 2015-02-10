package parc

type Grammar struct {
	slots map[string]*grammarSlot
	Start string
}

type MatchFunc func([]byte) (int, bool)

type grammarSlot struct {
	Type      slotType
	Alts      []*grammarSlot // for alternation
	Name      string         // for non-terminal
	Slot      *grammarSlot   // for non-terminal
	Continue  *grammarSlot
	MatchFunc MatchFunc // for terminal
	Id        string    // for debug
}

type slotType int

const (
	slotAlt slotType = iota
	slotNonTerminal
	slotTerminal
	slotReturn
	slotFinish
)
