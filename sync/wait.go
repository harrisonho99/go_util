package sync

import (
	"sync"
)

type blocker chan struct{}

type Waiter struct {
	atom     uint
	blockers []blocker
	mu       sync.Mutex
}

func (w *Waiter) Add(n uint) {
	w.atom += n
	var i uint
	for ; i < n; i++ {
		w.blockers = append(w.blockers, make(blocker))
	}
}

func (w *Waiter) Done() {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.atom == 0 {
		panic("Number of atom less than 0, you should call Enough Waiter.Add()")
	} else {
		w.atom--
		close(w.blockers[0])
		w.blockers = w.blockers[1:]
	}
}

func (w *Waiter) Wait() {
	for _, v := range w.blockers {
		for range v {
		}
	}
}
