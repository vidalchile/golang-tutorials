package main

import (
	"fmt"
)

func main() {
	x := 0
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++ // infinite loop without this
	}

	fmt.Print("-------------------------------------\n")

	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	}

	fmt.Print("-------------------------------------\n")

	names := []string{"mario", "luigi", "yoshi", "peach"}
	for i := 0; i < len(names); i++ {
		fmt.Println("value name is:", names[i])
	}

	fmt.Print("-------------------------------------\n")

	for index, value := range names {
		fmt.Printf("the value at pos %v is %v \n", index, value)
		value = "new string"
	}

	fmt.Print("-------------------------------------\n")

	for _, value := range names {
		fmt.Print(value, ",")
		value = "new string"
	}

	// changing val in a loop does not mutate the original slice
	fmt.Println(names)
}
