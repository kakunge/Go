package main

import (
	"fmt"
	"strings"
)

func main() {
	fn1, sn1 := getInitials("Hitori Gotoh")

	fmt.Println(fn1, sn1)
	
	fn2, sn2 := getInitials("Ryo Yamada")

	fmt.Println(fn2, sn2)

	fn3, sn3 := getInitials("Kita")

	fmt.Println(fn3, sn3)
}

func getInitials(m string) (string, string) {
	s := strings.ToUpper(m)
	names := strings.Split(s, " ")
	var initials []string

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}