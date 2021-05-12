package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"log"
)

const file = "perf_recursion_iteration.log"

// This `fact` function calls itself until it reaches the base case of `fact(0)`.
func fact(n int) float64 {
	if n == 0 {
		return 1
	}
	return float64(n) * fact(n-1)
}

func iter(n int) float64 {

	var res float64 = 1
	var i int

	if n == 0 {
		return 1
	}

	for i = 1; i <= n; i++ {
		res = res * float64(i)
	}

	return res

}

func main() {

	var n, j int
	var info string

	i, _ := strconv.Atoi(os.Args[1])
	n = int(i)

	proc := os.Args[2]

	start := time.Now()
	end := time.Since(start)

	switch proc {
	case "r":
		for j = 1; j < n; j++ {
			fmt.Println(fact(j))
		}
		fmt.Println(fact(j))

		end = time.Since(start)
		fmt.Println("Execution time recursion:", end)

		info =  "Recursion " + "n = " + strconv.Itoa(n) +  " # " + end.String()
	case "i":
		for j = 1; j < n; j++ {
			fmt.Println(iter(j))
		}	
		fmt.Println(iter(j))

		end = time.Since(start)	
		fmt.Println("Execution time iteration:", end)

		info =  "Iteration " + "n= " + strconv.Itoa(n) + " # " + end.String()
	}

	Log(file, info)
}

func Log(file string, info string) {

	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	d := []byte(info + "\n")
	if _, err := f.Write([]byte(d)); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}