package main

import (
	"testing"
)

const iterNum = 10

func BenchmarkEmptyAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := make([]int, 0)
		for j := 0; i < iterNum; j++ {
			data = append(data, j)
		}
	}
}
func BnechmarkPreallocAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := make([]int, 0, iterNum)
		for j := 0; j < iterNum; j++ {
			data = append(data, j)
		}
	}
}
