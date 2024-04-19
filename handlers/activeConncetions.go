package handlers

import "sync"

type activeConncetion struct {
	a      int
	locker sync.Mutex
}

func (a *activeConncetion) Add(x int) int {
	var tmp int
	a.locker.Lock()
	a.a += x
	tmp = a.a
	a.locker.Unlock()
	return tmp
}
