package main

import "fmt"

func main() {
	fmt.Println("[String]")

	var name1 string = "apple"
	var name2 = "banana"
	var name3 string

	fmt.Println(name1, name2, name3)
	fmt.Println("------------")

	name3 = "cherry"

	fmt.Println(name3)
	fmt.Println("------------")

	name4 := "durian"

	fmt.Println(name4)

	fmt.Println("\n[Integer]")

	var age1 int = 20
	var age2 = 15
	age3 := 19

	fmt.Println(age1, age2, age3)

	fmt.Println("\n[Float]")

	var score1 float32 = 25.99
	var score2 float64 = 438572359273735983892.32472
	score3 := 1.5

	fmt.Println(score1, score2, score3)
}