package main

import (
	"fmt"
)

func main() {
	mybill := newBill("apple")

	fmt.Println(mybill.format())
}