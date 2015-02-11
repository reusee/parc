package parc

type _Thread struct {
	slot     *GrammarSlot
	index    int
	stackTop *stackNode
}

type stackNode struct {
	slot   *GrammarSlot
	index  int
	parent *stackNode
}

var debug = false

func dumpThreads(threads []_Thread) {
	pt("--- threads ---\n")
	for _, thread := range threads {
		pt("%s %d\n", thread.slot.Name, thread.index)
		for e := thread.stackTop; e != nil; e = e.parent {
			id := e.slot.Name
			if len(id) == 0 {
				id = e.slot.Type.String()
			}
			pt("\t%v %d\n", id, e.index)
		}
	}
}

func (g *Grammar) Match(input []byte) bool {
	bottomNode := &stackNode{
		slot: &GrammarSlot{
			Type: slotFinish,
			Name: "Finish",
		},
	}
	threads := []_Thread{
		_Thread{
			slot:     g.slots[g.start],
			index:    0,
			stackTop: bottomNode,
		}}
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
					newThreads = append(newThreads, _Thread{
						slot:     alt,
						index:    thread.index,
						stackTop: thread.stackTop,
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
				thread.stackTop = &stackNode{
					slot:   thread.slot.Continue,
					index:  thread.index,
					parent: thread.stackTop,
				}
				thread.slot = slot
				newThreads = append(newThreads, thread)
			case slotReturn:
				node := thread.stackTop
				thread.stackTop = node.parent
				thread.slot = node.slot
				newThreads = append(newThreads, thread)
			case slotTerminal:
				if n, ok := thread.slot.MatchFunc(input[thread.index:]); ok {
					thread.index += n
					thread.slot = thread.slot.Continue
					newThreads = append(newThreads, thread)
				}
			case slotFinish:
				if thread.index == len(input) {
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
