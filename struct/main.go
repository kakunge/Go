package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
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

	switch opt {
	case "a":
		name, _ := getInput("Item name : ", reader)
		price, _ := getInput("Item price : ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}

		b.addItem(name, p)

		fmt.Println("item added -", name, price)

		promptOptions(b)
	case "s":
		b.save()

		fmt.Println("you choose save the bill", b.name)
	case "t":
		tip, _ := getInput("Tip amount ($) : ", reader)

		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}

		b.updateTip(t)

		fmt.Println("tip added -", tip)

		promptOptions(b)
	default:
		fmt.Println("wrong")
		promptOptions(b)
	}
}