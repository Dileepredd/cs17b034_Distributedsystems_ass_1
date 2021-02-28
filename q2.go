package main

import (
	"fmt"
	"sync"
)

const shape = 8

const buffersize = 40

const noofworkers = 40

var result int

var lock sync.Mutex

type config struct {
	col [shape]int
}

func check(curconfig *config) bool {
	// for each point
	for i := 0; i < shape; i++ {
		// check row
		row := curconfig.col[i]
		for j := 0; j < shape; j++ {
			if row == curconfig.col[j] && j != i {
				return false
			}
		}
		// check -1diagonal
		for di, dj := row, i; di >= 0 && dj >= 0; di, dj = di-1, dj-1 {
			if di == curconfig.col[dj] && dj != i {
				return false
			}
		}
		for di, dj := row, i; di < shape && dj < shape; di, dj = di+1, dj+1 {
			if di == curconfig.col[dj] && dj != i {
				return false
			}
		}
		// check +1diagonal
		for di, dj := row, i; di >= 0 && dj < shape; di, dj = di-1, dj+1 {
			if di == curconfig.col[dj] && dj != i {
				return false
			}
		}
		for di, dj := row, i; di < shape && dj >= 0; di, dj = di+1, dj-1 {
			if di == curconfig.col[dj] && dj != i {
				return false
			}
		}
	}
	fmt.Println(curconfig.col, "cleared")
	return true
}

func worker(id int, ch chan config, waitgroup *sync.WaitGroup) {
	for curconfig := range ch {
		if check(&curconfig) {
			lock.Lock()
			result++
			lock.Unlock()
		}
		waitgroup.Done()
	}
}

func generate(index int, configurations *config, ch chan config, waitgroup *sync.WaitGroup) {
	if index >= shape {
		waitgroup.Add(1)
		ch <- *configurations
		return
	}
	for i := 0; i < shape; i++ {
		configurations.col[index] = i
		generate(index+1, configurations, ch, waitgroup)
	}
}

func main() {
	result = 0
	var waitgroup sync.WaitGroup
	ch := make(chan config, buffersize)
	// start workers
	for i := 0; i < noofworkers; i++ {
		func() {
			go worker(i, ch, &waitgroup)
		}()
	}
	// genrate configurations
	var configurations config
	generate(0, &configurations, ch, &waitgroup)
	waitgroup.Wait()
	close(ch)
	fmt.Println("total number of possible configurations are :", result)
}
