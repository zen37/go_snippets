package main

import "fmt"

const (
	en       = "hello"
	jp       = "こんにちは"
	ch_simpl = "你好" // hello in chinese simplified

)

func main() {

	print(en)
	fmt.Println()
	print(jp)
	fmt.Println()
	print(ch_simpl)
	fmt.Println()

}

func print(text string) {

	fmt.Println(text)

	for i := 0; i < len(text); i++ {
		fmt.Print(text[i])
		fmt.Print(" ")
	}
	fmt.Println()

	for _, val := range text {
		fmt.Printf("%U ", val)
	}

	fmt.Println()

}
