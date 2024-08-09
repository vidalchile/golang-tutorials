package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
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

func main() {
	n, a := getInitials("cristian vidal")
	fmt.Println(n, a)

	n2, a2 := getInitials("cristian")
	fmt.Println(n2, a2)

	n3, a3 := getInitials("cristian vidal muñoz")
	fmt.Println(n3, a3)
}
