package main

import "fmt"

type pet struct {
	age int
	species string
	name string
}

func main() {
	fmt.Println(pet{5, "dog", "Astro"})

	fmt.Println(pet{name:"Astro", age: 5, species: "dog"})

	fmt.Println(pet{name:"Fred"})

	fmt.Println(&pet{9, "cat", "Little Peach"})

	s := pet{name: "Grey"}
	fmt.Println(s.name)

	sp := s
	fmt.Println(sp.name)

	sp.age = 7
	fmt.Println(sp.age)
}