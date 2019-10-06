package main

import (
	"fmt"
	"time"
	"sync"
)

func Work(wg *sync.WaitGroup, n int) {
	defer wg.Done()

	fmt.Println("start", n)
	time.Sleep(time.Second)
	fmt.Println("done", n)
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i <10; i++ {
		go Work(wg, i)
	}

	wg.Wait()
	fmt.Println("Main")
}

// start
// done
// Main