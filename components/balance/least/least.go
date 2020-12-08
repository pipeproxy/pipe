package least

import (
	"container/list"
	"math/rand"
	"sync"

	"github.com/pipeproxy/pipe/components/balance"
)

type element struct {
	index uint64
	size  uint64
}

type Least struct {
	list *list.List
	mut  sync.Mutex
}

func NewLeast() *Least {
	return &Least{
		list: list.New(),
	}
}

func (l *Least) Init(size uint64) {
	for uint64(l.list.Len()) < size {
		elem := &element{
			index: uint64(l.list.Len()),
		}
		if rand.Intn(2) == 0 {
			l.list.PushFront(elem)
		} else {
			l.list.PushBack(elem)
		}
	}
}

func (l *Least) InUse(fun func(i uint64)) {
	l.mut.Lock()
	target := l.list.Front()
	targetElem := target.Value.(*element)

	targetElem.size++
	next0 := target.Next()
	for next := next0; ; next = next.Next() {
		if next == nil {
			if next0 != next {
				l.list.MoveToBack(target)
			}
			break
		}
		nextElem := next.Value.(*element)
		if targetElem.size < nextElem.size {
			l.list.MoveBefore(target, next)
		}
	}
	index := targetElem.index
	l.mut.Unlock()

	fun(index)

	l.mut.Lock()
	targetElem.size--
	prev0 := target.Prev()
	for prev := prev0; ; prev = prev.Prev() {
		if prev == nil {
			if prev0 != prev {
				l.list.MoveToFront(target)
			}
			break
		}
		prevElem := prev.Value.(*element)
		if targetElem.size > prevElem.size {
			l.list.MoveAfter(target, prev)
		}
	}
	l.mut.Unlock()
}

func (l *Least) Clone() balance.Policy {
	c := NewLeast()
	c.Init(uint64(l.list.Len()))
	return c
}
