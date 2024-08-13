package main

import "fmt"

func main() {

	menu := map[string]float64{
		"soup":           10.1,
		"pie":            10.2,
		"salad":          10.3,
		"toffee pudding": 10.4,
	}

	fmt.Println(menu)

	fmt.Println(menu["salad"])

	// ints as key type
	phoneBook := map[int]string{
		267584967: "mario",
		984759373: "luigi",
		845775485: "peach",
	}

	for key, value := range phoneBook {
		fmt.Println(key, "-", value)
	}

	fmt.Println(phoneBook)
	fmt.Println(phoneBook[845775485])

	phoneBook[984759373] = "bowser"
	fmt.Println(phoneBook)

	phoneBook[647583927] = "yoshi"
	fmt.Println(phoneBook)
}
