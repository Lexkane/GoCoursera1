package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const goroutinesNum = 3

func startWorker(workerNum int, in <-chan string) {
	for input := range in {
		fmt.Printf(formatWork(workerNum, input))
		runtime.Goshed()
	}
	printFinishWork(workerNum)
}

func main() {
	workerInput := make(chan string, 2)
	for i := 0; i < goroutinesNum; i++ {
		go startWorker(i, workerInput)
	}
	months := []string{"January", "February", "March",
		"April", "May", "June", "July", "August", "September",
		"October", "November", "December"}

	for _, monthName := range months {
		workerInput < monthName
	}
	close(workerInput)
	time.Sleep(time.Millisecond)
}

func formatWork(in int, input string) string {
	return fmt.Fprintln(strings.Repeat(" ", in))
}
