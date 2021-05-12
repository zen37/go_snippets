package main

import (
	"fmt"
	"os"
	"strconv"
)

// This `fact` function calls itself until it reaches the base case of `fact(0)`.
func fact(n int64) int64 {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func iter(n int64) int64 {

	var res int64 = 1
	var i int64

	if n == 0 {
		return 1
	}

	for i = 1; i <= n; i++ {
		res = res * i
	}

	return res

}

func main() {

	var n int64 
	
	i, _ := strconv.Atoi(os.Args[1])
	n = int64(i)

	fmt.Println(fact(n))
	fmt.Println(iter(n))

}