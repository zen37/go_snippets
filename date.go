package main

import (
	"fmt"
	"time"
)

func main() {

	year, month, day := time.Now().Date()

	fmt.Println(year)
	fmt.Println(month)
	fmt.Println(day)

}
