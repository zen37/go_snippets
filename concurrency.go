package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

const thread int = 10

var x, i float64

const n float64 = 10000000 //2147483647 //12, 9223372036854775807

//var s []float64

//var m map[string]float64

func sqrt(i float64, j float64, wg *sync.WaitGroup) {

	defer wg.Done()

	//println("start =", int(i))
	//println("end =", int(j))

	for k := i; k <= j; k++ {
		x = math.Sqrt(k)
		//	fmt.Println(k, " - ", x)
	}

	//s = append(s, x)
	//m[strconv.Itoa(int(i))+"-"+strconv.Itoa(int(j))] = x

}

func main() {

	var wg sync.WaitGroup

	//m = make(map[string]float64)

	start := time.Now()

	for j := 1; j <= thread; j++ {

		wg.Add(1)

		s := (j-1)*int(n)/thread + 1
		e := j * int(n) / thread

		go sqrt(float64(s), float64(e), &wg)
	}

	wg.Wait()

	fmt.Println("GO duration: ", time.Since(start).Milliseconds(), "ms")

	//sort.Float64s(s)
	//fmt.Println(s)
	/*
		fmt.Println("map:", m)

		for key, value := range m {
			fmt.Println(key, value)
		}
	*/

}
