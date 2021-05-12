package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()

	day := t.Weekday()
	fmt.Println(t)
	fmt.Println(day)

	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}
