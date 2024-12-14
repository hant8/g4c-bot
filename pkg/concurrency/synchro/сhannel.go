package synchro

import "fmt"

func workItemChannel(done chan bool) {
	fmt.Println("Working...")
	done <- true
}

func RunChannel() {
	done := make(chan bool)
	go workItemChannel(done)
	<-done
	fmt.Println("Work completed.")
}
