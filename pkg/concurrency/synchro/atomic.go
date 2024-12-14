package synchro

import (
	"fmt"
	"sync/atomic"
	"time"
)

func RunAtomic() {
	var counter int32

	for i := 0; i < 10; i++ {
		go atomic.AddInt32(&counter, 10)
	}

	time.Sleep(time.Second)
	fmt.Println("Counter:", counter)
}
