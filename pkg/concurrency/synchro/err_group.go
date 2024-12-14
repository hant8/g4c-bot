package synchro

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func RunErrGroup() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		fmt.Println("Task 1 started")
		time.Sleep(2 * time.Second)
		fmt.Println("Task 1 finished")
		return nil
	})

	g.Go(func() error {
		fmt.Println("Task 2 started")
		time.Sleep(1 * time.Second)
		return errors.New("task 2 failed")
	})

	g.Go(func() error {
		fmt.Println("Task 3 started")
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("Task 3 finished")
			return nil
		case <-ctx.Done():
			fmt.Println("Task 3 cancelled due to error")
			return ctx.Err()
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("All tasks finished successfully")
	}
}
