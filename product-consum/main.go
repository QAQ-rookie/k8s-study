package main

import (
	"fmt"
	"runtime"
	"sync"
)

func producer(ch chan<- int, n int, wg *sync.WaitGroup) {
	defer func() {
		close(ch)
		wg.Done()
	}()
	for i := 0; i < n; i++ {
		ch <- i
	}
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	for {
		select {
		case data, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(data)
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(2)
	go producer(ch, 10, wg)
	go consumer(ch, wg)
	wg.Wait()
	runtime.GOMAXPROCS(2)
}
