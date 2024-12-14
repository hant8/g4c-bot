package synchro

import (
	"sync"
	"time"
)

var ready = false

func workItemCond(cond *sync.Cond, wg *sync.WaitGroup) {
	defer wg.Done()
	cond.L.Lock()
	for !ready {
		cond.Wait()
	}
	cond.L.Unlock()
}

func RunCond() {
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup

	wg.Add(1)
	go workItemCond(cond, &wg)
	time.Sleep(3 * time.Second)

	cond.L.Lock()
	ready = true
	cond.Signal()
	cond.L.Unlock()

	wg.Wait()
}
