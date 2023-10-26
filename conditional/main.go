package main

import (
	"fmt"
)

func main() {
	age := 45

	fmt.Println(age <= 70)
	fmt.Println(age >= 70)
	fmt.Println(age == 45)
	fmt.Println(age != 70)

	if age < 30 {
		fmt.Println("less than 30")
	} else if age < 40 {
		fmt.Println("less than 40")
	} else {
		fmt.Println("not less than 40")
	}

	fmt.Println("------------")

	fruits := []string{"apple", "banana", "cherry", "durian"}

	for index, value := range fruits {
		if index == 1 {
			fmt.Println("continue", index)
			continue
		}
		if index > 2 {
			fmt.Println("break", index)
			break
		}

		fmt.Printf("%v %v\n", index, value)
	}
}