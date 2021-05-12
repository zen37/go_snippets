package main

import (
	"fmt"
	"time"
)

type err struct {
	arg int
	txt string
}

func (e err) Error() string {
	return fmt.Sprintf("%d / %s", e.arg, e.txt)
}

func f(arg int) (int, error) {
	if arg != 42 {
		return -1, err{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	start := time.Now()

	for i := 1; i < 100000; i++ {
		_, e := f(i)
		if _, ok := e.(err); ok {
			// fmt.Println(r.arg)
			// fmt.Println(r.txt)
		}
	}

	fmt.Println("Execution time: ", time.Since(start).Nanoseconds(),"ns")

}