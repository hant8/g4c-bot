package synchro

import (
	"fmt"
	"sync"
)

func RunSyncPool() {
	pool := sync.Pool{
		New: func() any {
			return "default"
		},
	}

	pool.Put(true)
	pool.Put("item1")
	pool.Put(false)
	pool.Put("item2")
	pool.Put(true)

	fmt.Println(pool.Get().(bool))
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
}
