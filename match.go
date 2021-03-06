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
	newThreads := []_Thread{}
	uniqueThreads := make(map[_Thread]struct{})
	addThread := func(thread _Thread) {
		if _, ok := uniqueThreads[thread]; !ok {
			uniqueThreads[thread] = struct{}{}
			newThreads = append(newThreads, thread)
		}
	}
	for len(threads) > 0 {
		if debug {
			dumpThreads(threads)
			pt("\n\n")
		}
		for _, thread := range threads {
			switch thread.slot.Type {
			case slotAlt:
				for _, alt := range thread.slot.Alts {
					addThread(_Thread{
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
				// push
				thread.stackTop = &stackNode{
					slot:   thread.slot.Continue,
					index:  thread.index,
					parent: thread.stackTop,
				}
				thread.slot = slot
				addThread(thread)
			case slotReturn:
				// pop
				node := thread.stackTop
				thread.stackTop = node.parent
				thread.slot = node.slot
				addThread(thread)
			case slotTerminal:
				if n, ok := thread.slot.MatchFunc(input[thread.index:]); ok {
					thread.index += n
					thread.slot = thread.slot.Continue
					addThread(thread)
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
