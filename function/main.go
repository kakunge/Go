package main

import (
	"fmt"
	"math"
)

func main() {
	sayHello("Alice")
	sayBye("Alice")
	sayHello("Bob")

	fmt.Println("------------")

	cycleNames([]string{"Carol", "Dave", "Eve"}, sayHello)
	cycleNames([]string{"Carol", "Dave", "Eve"}, sayBye)

	fmt.Println("------------")

	area1 := circleArea(10.5)
	area2 := circleArea(15)

	fmt.Println(area1, area2)
}

func sayHello(m string) {
	fmt.Printf("Hi %v\n", m)
}

func sayBye(m string) {
	fmt.Printf("Bye %v\n", m)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}