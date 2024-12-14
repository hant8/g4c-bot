package synchro

import (
	"fmt"
	"time"
)

func workItemSemaphore(id int, semaphore chan struct{}) {
	semaphore <- struct{}{}
	fmt.Printf("Worker %d is working\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Worker %d done\n", id)
	<-semaphore
}

func RunSemaphore() {
	const maxWorkers = 3
	semaphore := make(chan struct{}, maxWorkers)

	for i := 0; i < 10; i++ {
		go workItemSemaphore(i, semaphore)
	}

	time.Sleep(10 * time.Second)
}
