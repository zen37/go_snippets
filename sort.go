package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
	"flag"
	"strconv"
	"log"
)

var arr []int
var algo *string
var size int


func init() {
	
	var err error
	const default_size = 11 

	algo = flag.String("algo", "s", "algorithm used for processing")
	flag.Parse()

	if len(flag.Args()) == 0 {
		size = default_size
	} else {
		size, err = strconv.Atoi(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
	}


}

func create_arr() {

	for i := 0; i < size; i++ {
		arr = append(arr, rand.Intn(100))
	}

}

func sort_standard(arr []int) {

	fmt.Println("sort.Slice before: ", arr)

	start:= time.Now()

	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	fmt.Println("Standard sort.Slice execution time: ", time.Since(start).Nanoseconds(),"ns")
	
	fmt.Println("sort.Slice after: ", arr)
	fmt.Println()

}

func sorting(arr []int) {

	fmt.Println("sort.Slice before: ", arr)

	start:= time.Now()

	sort.Ints(arr)

	fmt.Println("sort.Ints execution time: ", time.Since(start).Nanoseconds(),"ns")
	
	fmt.Println("sort.Slice after: ", arr)
	fmt.Println()

}


func iter(arr []int) {

	var key, i int //k

	fmt.Println("iteration before: ", arr)

	start := time.Now()

	n := len(arr)

	for j := 1; j <= n-1; j++ {
		key = arr[j]

		i = j - 1

		for i >= 0 && arr[i] > key {
		//	k = k + 1
			arr[i+1] = arr[i]
			i = i - 1
			arr[i+1] = key

		}
	}
	fmt.Println("Iteration execution time: ", time.Since(start).Nanoseconds(),"ns")

	fmt.Println("iteration after: ", arr)
	//fmt.Println()
	//fmt.Println("how many times with iteration: ", k)
	fmt.Println()

}

func main() {

	create_arr()

	arr1 := make([]int, len(arr))
	arr2 := make([]int, len(arr))
	arr3 := make([]int, len(arr))

	copy(arr1, arr)
	copy(arr2, arr)
	copy(arr3, arr)

	switch *algo {
	case "c": //custom
		iter(arr1)
	case "s": //standard
		sort_standard(arr2)
	case "b": //both
		iter(arr1)	
		sort_standard(arr2)	
		sorting(arr3)
	}
//	fmt.Println("original end: ", arr)
}