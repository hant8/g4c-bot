package synchro

import (
	"fmt"
	"sync"
	"time"
)

func workItemSBarrier(id int, startBarrier *sync.WaitGroup) {
	startBarrier.Wait()
	fmt.Printf("Worker %d started\n", id)
}

func RunBarrier() {
	var startBarrier sync.WaitGroup
	startBarrier.Add(1)

	for i := 1; i <= 5; i++ {
		go workItemSBarrier(i, &startBarrier)
	}

	fmt.Println("Releasing workers")
	startBarrier.Done()
	time.Sleep(1 * time.Second)
}
