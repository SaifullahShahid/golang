package main

import (
	"fmt"
	"time"
)

func main() {
	var n = 1000000000 // 1 billion only testslice2 able to handle
	//var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	//fmt.Printf("Total time without preallocation(testSlice): %v \n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation(testSlice2): %v \n", timeLoop(testSlice2, n))

}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for i := 0; i < n; i++ {
		slice = append(slice, i)
	}
	return time.Since(t0)
}
