package synchro

import (
	"context"
	"fmt"
	"time"
)

func workItemContext(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopped\n", id)
			return
		default:
			fmt.Printf("Worker %d is working\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func RunContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for i := 1; i <= 3; i++ {
		go workItemContext(ctx, i)
	}

	time.Sleep(5 * time.Second)
}
