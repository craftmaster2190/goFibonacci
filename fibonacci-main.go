package main

import (
	"fmt"
	"goFibonacci/lib"
)

func main() {
	var i uint64 = 0
	for i < 20 {
		fmt.Printf("Fibonacci(%d) = %d\n", i, lib.Fibonacci(i))
		i++
	}
}
