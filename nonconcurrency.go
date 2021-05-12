package main

import (
	"fmt"
	"math"
	"time"
)

const thread int = 10

var x, i float64

const n float64 = 10000000 //2147483647 //12, 9223372036854775807

func sqrt(i float64, j float64) {

	for k := i; k <= j; k++ {
		x = math.Sqrt(k)
		//fmt.Println(k, " - ", x)
	}

}

func main() {

	start := time.Now()

	for j := 1; j <= thread; j++ {

		s := (j-1)*int(n)/thread + 1
		e := j * int(n) / thread

		sqrt(float64(s), float64(e))
	}

	fmt.Println("GO duration: ", time.Since(start).Milliseconds(), "ms")

}
