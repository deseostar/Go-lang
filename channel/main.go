package main

import (
	"fmt"
	"time"
	"sync"
)

// func Work(wg *sync.WaitGroup, n int) {
// 	defer wg.Done()

// 	fmt.Println("start", n)
// 	time.Sleep(time.Second)
// 	fmt.Println("done", n)
// }

type Worker struct {
	Name string
}

func (w Worker) Work(c chan int, wg *sync.WaitGroup) {
	for task := range c {
		// task := <- c

		time.Sleep(time.Second)
		fmt.Printf("%s: Done[%d]\n", w.Name, task)
		wg.Done()
	}
	fmt.Println("Stop:", w.Name)
}

// func Send(c chan int, val int) {
// 	c <- val
// }

func main() {
	// c := make(chan int)
	// // c := make(chan int, 10) 버퍼 넣으면 블락이 안걸림

	// go Send(c, 100)  //send가 블락 되어있어도 go 쓰면 됨

	// fmt.Println(<-c)

	w := Worker{
		Name: "Worker_1",
	}
	workers := make([]Worker, 0, 10)
	wg := &sync.WaitGroup{}
	c := make(chan int)

	for i := 0; i <10; i++ {
		workers = append(workers, Worker{
			Name: fmt.Sprintf("Worker_%d", i),
		})
	}

	for _, worker := range workers {
		go worker.Work(c, wg)
	}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		c <- i
	}
	close(c)
	wg.Wait()

	val, ok := <- chan c
	fmt.Println(val, ok)


	// go w.Work(c)
	
	// c <- 1
	// c <- 2
	// c <- 3
	// c <- 4
	// c <- 5
}

// start
// done
// Main