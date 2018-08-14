package main

import (
	"log"
	"sync"
)

var counter int
var l sync.Mutex

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.Lock()
			counter++
			l.Unlock()
		}()
	}

	wg.Wait()
	log.Println(counter)
}
