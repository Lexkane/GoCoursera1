package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	iterationsNum = 7
	gorouinesNum  = 3
)

func startWorker(in int, wg *sync.Waitgroup) {
	defer wg.Done()
	for j := 0; j < iterationsNum; j++ {
		fmt.Printf(formatWork(in, j))
		rutime.GoShed()
	}
}

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < gorouinesNum; i++ {
		wg.Add(1)
		go startWorker(i, wg)
	}
	time.Sleep(time.Millisecond)
	//fmt.Scaln()
	wg.Wait()
}
