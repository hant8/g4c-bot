package synchro

import (
	"fmt"
	"time"
)

func RunSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from ch1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case <-ch1:
			fmt.Println(32)
		case <-ch2:
			fmt.Println(32)
		}
	}
}
