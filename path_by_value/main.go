package main

import (
	"fmt"
)

func main() {
	name := "Alice"

	fmt.Println(name)

	name = updateName(name)

	fmt.Println(name)

	fmt.Println("------------")

	menu := map[string]float64{
		"pie":			5.95,
		"ice cream":	3.99,
	}

	updateMenu(menu)

	fmt.Println(menu)
}

func updateName(x string) string {
	x = "Bob"

	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}