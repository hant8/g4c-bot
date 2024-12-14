package synchro

import (
	"fmt"
	"sync"
)

var (
	counterRW int
	rwMutex   sync.RWMutex
)

func read(wg *sync.WaitGroup) {
	defer wg.Done()
	rwMutex.RLock()
	fmt.Println("Reading Counter:", counterRW)
	rwMutex.RUnlock()
}

func write(wg *sync.WaitGroup) {
	defer wg.Done()
	rwMutex.Lock()
	counterRW++
	rwMutex.Unlock()
}

func RunRWMutex() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go read(&wg)
		wg.Add(1)
		go write(&wg)
	}

	wg.Wait()
}
