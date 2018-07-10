package main

import (
	"errors"
	"log"
	"sync"
	"time"
)

func syncDemo() {
	log.Println("start", time.Now())

	var t1, t2 int
	var wg sync.WaitGroup
	wg.Add(2)
	go func(id int, wg *sync.WaitGroup) {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		log.Println("go 1 ", id, time.Now())
		t1 = id + 3
	}(1, &wg)

	go func(id int, wg *sync.WaitGroup) {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		log.Println("go 2 ", id, time.Now())
		t2 = id + 3
	}(2, &wg)

	wg.Wait()

	time.Sleep(time.Second * 5)
	log.Println("done", time.Now(), t1, t2)
}

func channelDemo() {
	log.Println("start", time.Now())

	t1 := make(chan int)
	t2 := make(chan int)
	err := make(chan error)

	go func(id int) {
		time.Sleep(3 * time.Second)
		log.Println("go 1 ", id, time.Now())
		t1 <- 0
		err <- errors.New("go 1 error")
		// close(t1)
	}(1)

	go func(id int) {
		time.Sleep(1 * time.Second)
		log.Println("go 2 ", id, time.Now())
		t2 <- id + 3

	}(2)

	time.Sleep(time.Second * 5)
	// log.Println("done", time.Now(), <-t1, <-t2)

	for i := 0; i < 3; i++ {
		select {
		case n := <-err:
			log.Println(n)
		case total := <-t1:
			log.Println("total", total)
		case size := <-t2:
			log.Println("size", size)
		}
	}

}

func main() {
	channelDemo()
}
