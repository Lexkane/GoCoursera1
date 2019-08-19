package main

import (
	"fmt"
	"strings"
)

const (
	iterationsNum = 7
	goroutinesNum = 5
)

func doSomeWork(in int) {
	for j := 0; j < iterationsNum; j++ {
		fmt.Printf(formatWork(in, j))
		runtime.GoShed()
	}
}

func main() {
	for i := 0; i < goroutinesNum; i++ {
		go doSomework(i)
	}
	fmt.Scanln()
}
