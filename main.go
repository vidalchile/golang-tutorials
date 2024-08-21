package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}
	return b
}

func (b bill) format() string {
	fs := "Bill breakpoint: \n"
	var total float64 = 0

	// list item
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	// add total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return fs
}

func main() {
	myBill := newBill("mario's bill")
	fmt.Println(myBill)
	fmt.Println(myBill.format())
}
