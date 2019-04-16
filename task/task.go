package task

import (
	"sync"
)

func Spawn(wg *sync.WaitGroup, f func()) {

	if wg != nil {
		wg.Add(1)
	}

	go func() {
		if wg != nil {
			defer wg.Done()
		}

		defer func() {
			if p := recover(); p != nil {
				//TODO extra logging
				panic(p)
			}
		}()
		f()

	}()
}
