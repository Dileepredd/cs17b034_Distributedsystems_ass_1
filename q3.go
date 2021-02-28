package main

import (
	"fmt"
	"strconv"
	"sync"
)

func print(str string, ch chan int, waitgroup *sync.WaitGroup) {
	ch <- 1
	fmt.Println("worker id:", str, "number of workers in critical section:", len(ch), "out of:", cap(ch))
	// time.Sleep((time.Second))
	<-ch
	waitgroup.Done()
}

func main() {
	ch := make(chan int, 20)
	var waitgroup sync.WaitGroup
	waitgroup.Add(100)
	for i := 0; i < 100; i++ {
		func() {
			go print(strconv.Itoa(i), ch, &waitgroup)
		}()
	}
	waitgroup.Wait()
}
