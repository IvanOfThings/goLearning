package main

import (
	"runtime"
	"sync/atomic"
)

type TicketStore struct {
	ticket *uint64
	done   *uint64
	slots  []string
}

func (ts *TicketStore) Put(s string) {
	t := atomic.AddUint64(ts.ticket, 1) - 1 //draw a ticker
	ts.slots[t] = s
	for !atomic.CompareAndSwapUint64(ts.done, t, t+1) {
		runtime.Gosched()
	}

}

func (ts *TicketStore) GetDone() []string {
	return ts.slots[:atomic.LoadUint64(ts.done)+1]
}
