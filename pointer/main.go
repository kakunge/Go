package main

import (
	"fmt"
)

func main() {
	name := "Alice"

	fmt.Println(name)

	fmt.Println("Memory address of name :", &name)

	m := &name

	fmt.Println("m :", m)
	fmt.Println("value :", *m)

	updateName(m)

	fmt.Println(name)
}

func updateName(x *string) {
	*x = "Bob"
}