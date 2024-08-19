package main

import "fmt"

func updateName(n string) {
	// no actualizamos la variable name
	// actualizamos la copia de la variable name
	n = "cristian"
}

func updateNameWithReturn(x string) string {
	// no actualizamos la variable name
	// actualizamos la copia de la variable name
	x = "cristian"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 1.3
}

func main() {
	// group A types -> strings, ints, bools, floats, arrays, structs (non-pointer values)

	name := "carolina"

	// al pasar una variable a una funcion go crea una copia
	updateName(name)

	fmt.Println(name)

	// al crear una variable go la almacena en un bloque de memoria
	/*

		|----name----|------n-------|
		|  0x001     |  0x002  	    |
		|------------|--------------|
		| "carolina" | "cristian"   |
		|------------|--------------|

	*/

	// ¿Como solucionamos esto?
	name = updateNameWithReturn(name)
	println(name)

	// group B types -> slices, maps, functions
	menu := map[string]float64{
		"pie":       1.1,
		"ice cream": 1.2,
	}

	// para este tipo de caso si es posible modificar el valor original (pointer wrapper values)
	updateMenu(menu)
	fmt.Println(menu)
}
