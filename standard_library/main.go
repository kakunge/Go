package main

import (
	"fmt"
	"strings"
	"sort"
)

func main() {
	aisatsu := "ohayou konnichiha konbanha"

	fmt.Println(aisatsu)
	fmt.Println(strings.Contains(aisatsu, "konnichiha"))
	fmt.Println(strings.Contains(aisatsu, "ohayou!"))
	fmt.Println(strings.ReplaceAll(aisatsu, "konnichiha", "sayonara"))
	fmt.Println(aisatsu)
	fmt.Println(strings.ToUpper(aisatsu))
	fmt.Println(strings.Index(aisatsu, "kon"))
	fmt.Println(strings.Split(aisatsu, " "))
	fmt.Println("------------")

	numbers := []int{3, 6, 2, 5, 8, 1}

	fmt.Println(numbers)
	sort.Ints(numbers)
	fmt.Println(numbers)

	index := sort.SearchInts(numbers, 5)

	fmt.Println(index)

	index = sort.SearchInts(numbers, 7)

	fmt.Println(index)
}