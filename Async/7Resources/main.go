package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	iteratationsNum = 6
	goroutinesNum   = 5
	quotaLimit      = 2
)

func startWorker(int int, wg *sync.WaitGroup, quoaaCh chan struct{}) {
	quotaCh <- struct{}{}
	defer wg.Done()
	for j := 0; j < iteratationsNum; j++ {
		fmt.Printf(formatWork(in, j))
		//if j%2==0 {
		//<-quotaCh
		//quotaCh<-struct{}{}
		//}
		runtime.Gosched()
	}
	<-quotaCh
}

func main() {
	wg := &sync.WaitGroup{}
	quotaCh := make(chan struct{}, quotaLimit)
	for i := 0; i < goroutinesNum; i++ {
		wg.Add(1)
		go startWorker(i, wg, quotach)
	}
	time.Sleep(time.Millisecond)
	wg.Wait()
}

func formatWork(in, j int) string {
	return fmt.Sprintln(strings.Repeat(" ", in))
}
