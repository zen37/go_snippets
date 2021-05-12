package main

import (
	"fmt"
)

func main() {

	c := make(chan string)
	go count("sheep", c)
	msg := <-c
	fmt.Println(msg)
}

func count(t string, c chan string) {

	for i := 1; i < 5; i++ {
		fmt.Println("i=", i)
		//time.Sleep(time.Second * 1)
		c <- t

	}

	close(c)

}
