package main

import (
	"flag"
	"fmt"
)

var err_show *bool

func init() {

	err_show = flag.Bool("error", false, "this is a boolean for printing or not errors. default is false")
	fmt.Println("err_show before parse:", *err_show)

	flag.Parse()

	fmt.Println("err_show after parse:", *err_show)
}

func main() {

	fmt.Println("err_show main():", *err_show)

	if *err_show {
		fmt.Println("errors are printed")
	} else {
		fmt.Println("errors are not printed")
	}
}