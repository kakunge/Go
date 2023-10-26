package main

import (
	"fmt"
)

func main() {
	menu := map[string]float64{
		"soup":				4.99,
		"pie":				7.99,
		"salad":			6.99,
		"toffee pudding":	3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])
	
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}
	
	fmt.Println("------------")
	
	phonebook := map[int]string{
		12345678: "Alice",
		48324740: "Bob",
		72845279: "Carol",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[48324740])
	
	phonebook[72845279] = "Chuck"
	fmt.Println(phonebook)
	
	phonebook[48324740] = "Dave"
	fmt.Println(phonebook)
}