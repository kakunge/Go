package main

import (
	"fmt"	
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("------------")

	fruits := []string{"apple", "banana", "cherry", "durian"}

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	for index, value := range fruits {
		fmt.Println(index, value)
	}
}