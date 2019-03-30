package main

import (
	"runtime"
	"sync/atomic"
)

type Spinlock struct {
	state *int32
}

const free = int32(0)

func (l *Spinlock) Lock() {
	for !atomic.CompareAndSwapInt32(l.state, free, 42) { // 42 or any other value but 0 (each consumer should have his own value)
		runtime.Gosched() // Poker the scheduler
	}
}

func (l *Spinlock) Unlock() {
	atomic.StoreInt32(l.state, free) // Once atomic, always atomic
}
