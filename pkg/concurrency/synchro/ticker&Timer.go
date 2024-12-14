package synchro

import (
	"fmt"
	"time"
)

func RunTimer() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("Waiting...")
	_, ok := <-timer.C
	fmt.Println("Timer expired", ok)
}

func RunTicker() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for i := 0; i < 5; i++ {
		fmt.Println("Tick at:", <-ticker.C)
	}
}
