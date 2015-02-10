package parc

type stackEntry struct {
	slot  *GrammarSlot
	index int
}
type _Thread struct {
	slot  *GrammarSlot
	index int
	stack []stackEntry
}

var debug = false

func dumpThreads(threads []_Thread) {
	pt("--- threads ---\n")
	for _, thread := range threads {
		pt("%s %d\n", thread.slot.Name, thread.index)
		for i := len(thread.stack) - 1; i >= 0; i-- {
			e := thread.stack[i]
			pt("\t%v %d\n", e.slot.Name, e.index)
		}
	}
}

func (g *Grammar) Match(input []byte) bool {
	threads := []_Thread{
		{g.slots[g.start], 0, []stackEntry{
			stackEntry{
				slot: &GrammarSlot{
					Type: slotFinish,
					Name: "Finish",
				}}}},
	}
	var newThreads []_Thread
	for len(threads) > 0 {
		if debug {
			dumpThreads(threads)
			pt("\n\n")
		}
		for _, thread := range threads {
			switch thread.slot.Type {
			case slotAlt:
				for _, alt := range thread.slot.Alts {
					stack := make([]stackEntry, len(thread.stack))
					copy(stack, thread.stack)
					newThreads = append(newThreads, _Thread{
						slot:  alt,
						index: thread.index,
						stack: stack,
					})
				}
			case slotNonTerminal:
				slot := thread.slot.Slot
				if slot == nil {
					slot = g.slots[thread.slot.Symbol]
					thread.slot.Slot = slot
				}
				if slot == nil {
					panic("non-exists Non-terminal " + thread.slot.Symbol)
				}
				thread.stack = append(thread.stack, stackEntry{
					slot:  thread.slot.Continue,
					index: thread.index,
				})
				thread.slot = slot
				newThreads = append(newThreads, thread)
			case slotReturn:
				entry := thread.stack[len(thread.stack)-1]
				thread.stack = thread.stack[:len(thread.stack)-1]
				thread.slot = entry.slot
				newThreads = append(newThreads, thread)
			case slotTerminal:
				if n, ok := thread.slot.MatchFunc(input[thread.index:]); ok {
					thread.index += n
					thread.slot = thread.slot.Continue
					newThreads = append(newThreads, thread)
				}
			case slotFinish:
				if thread.index == len(input) && len(thread.stack) == 0 {
					return true
				}
			default:
				panic("not handled slot type " + thread.slot.Type.String())
			}
		}
		threads, newThreads = newThreads, threads
		newThreads = newThreads[0:0]
	}
	return false
}
