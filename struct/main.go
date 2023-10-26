package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	// mybill := newBill("Alice")

	// mybill.addItem("onion soup", 4.50)
	// mybill.addItem("veg pie", 8.95)
	// mybill.addItem("toffee pudding", 4.95)
	// mybill.addItem("coffee", 3.25)
	// mybill.updateTip(10)

	// fmt.Println(mybill.format())

	mybill := createBill()

	promptOptions(mybill)

	fmt.Println(mybill)
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create new bill name : ")
	
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create new bill name : ", reader)
	b := newBill(name)

	fmt.Println("Created the bill -", b.name)

	return b
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)

	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip) : ", reader)

	fmt.Println(opt)
}