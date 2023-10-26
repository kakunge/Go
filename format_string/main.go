package main

import "fmt"

func main() {
	month := "October"
	day := 26
	
	fmt.Println("Today is", month, day, "th")
	fmt.Printf("Today is %v %v th\n", month, day)
	fmt.Printf("Today is %q %q th\n", month, day)

	fmt.Printf("Type of day is %T\n", day)

	fmt.Printf("Float type : %f\n", 199.98)
	fmt.Printf("Float type : %0.2f\n", 199.98)

	var str = fmt.Sprintf("Today is %v %v th\n", month, day)
	fmt.Print(str)
}