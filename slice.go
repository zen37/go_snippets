package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

func proc(s []string) {
	fmt.Println("function started")

	// j := 0
	// for range s {
	// 	fmt.Println(s[j])
	// 	s[1] = "jdahjads"
	// 	j++
	// }

	// for i := 0; i < len(s); i++ {
	// 	s[1] = "what's that"
	// 	fmt.Println(i, " ", s[i])
	// }

	for i, e := range s {

		fmt.Println(i, " ", e)
		s[1] = "what's that"
		fmt.Println(i, " ", e)
	}
	fmt.Println("function ended")
}

func main() {

	people := []person{}
	people = append(people, person{Name: "Vera", Age: 24})
	people = append(people, person{Name: "John", Age: 88})
	people = append(people, person{
		Name: "Bob",
		Age:  111,
	})
	people = append(people, person{
		Name: "Gopher",
		Age:  2,
	})

	//sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	//fmt.Printf("%v", people)

	s := []string{"a", "b", "c"}

	fmt.Println("before passing slice: ", s)

	proc(s)

	fmt.Println("after passing slice: ", s)

	sort.Slice(people, func(i, j int) bool { return people[i].Age > people[j].Age })
	//fmt.Println("By age:", people)

	// for _, elem := range people {
	// 	fmt.Println(elem.Name, elem.Age)
	// }

}
