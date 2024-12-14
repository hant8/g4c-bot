package synchro

import (
	"fmt"
	"sync"
)

func RunSyncMap() {
	var sm sync.Map

	sm.Store("key1", "value1")
	sm.Store("key2", "value2")

	value, ok := sm.Load("key1")
	if ok {
		fmt.Println("Key1:", value)
	}

	sm.Range(func(key, value interface{}) bool {
		fmt.Printf("%s: %s\n", key, value)
		return true
	})
}
