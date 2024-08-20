package main

import "fmt"

func updateName(n string) {
	n = "cristian"
	// & gets the memory address of the value (pointer)
	fmt.Println("memory address of n is:", &n)
	fmt.Println("value n:", n)
}

func updatePointer(n *string) {
	*n = "kira"
}

func main() {
	// group A types -> strings, ints, bools, floats, arrays, structs (non-pointer values)

	name := "carolina"

	updateName(name)

	// & gets the memory address of the value (pointer)
	fmt.Println("memory address of name is:", &name)

	// * gets the value at the specified memory address
	m := &name // ahora es un puntero
	fmt.Println("memory address m:", m)
	fmt.Println("value at memory address m:", *m)

	fmt.Println(name)

	updatePointer(m)

	fmt.Println(name)

	/*

		|--name-------|----m----|
		|  0x001  	  |  0x002  |
		|-------------|---------|
		| "carolina"  | p0x001  |
		|-------------|---------|

	*/
}
