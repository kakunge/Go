package main

import "fmt"

func main() {
	fmt.Println("[Array]")

	var ages [3]int = [3]int{20, 25, 30}
	names := [4]string{"apple", "banana", "cherry", "durian"}

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	names[1] = "blueberry"

	fmt.Println(names, len(names))

	fmt.Println("\n[Slice]")

	var scores = []int{100, 99, 50}

	fmt.Println(scores, len(scores))

	scores[1] = 35

	fmt.Println(scores, len(scores))

	scores = append(scores, 77)

	fmt.Println(scores, len(scores))

	fmt.Println("\n[Range]")

	range1 := names[1:3]
	range2 := names[2:]

	fmt.Println(range1, range2)
}