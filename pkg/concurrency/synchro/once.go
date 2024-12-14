package synchro

import (
	"fmt"
	"sync"
)

func RunOnce() {
	var once sync.Once

	initializer := func() {
		fmt.Println("Initialized once.")
	}

	for i := 0; i < 3; i++ {
		go once.Do(initializer)
	}
}
