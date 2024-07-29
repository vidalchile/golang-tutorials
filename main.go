package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string) {
	fmt.Printf("Good morning %v \n", name)
}

func sayBye(name string) {
	fmt.Printf("Goodbye %v \n", name)
}

func cycleNames(names []string, f func(string)) {
	for _, name := range names {
		f(name)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	names := []string{"cristian", "carolina"}

	sayGreeting(names[0])
	sayGreeting(names[1])

	sayBye(names[0])
	sayBye(names[1])

	cycleNames(names, sayGreeting)
	cycleNames(names, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("circle 1 area is %0.3f & circle 2 area is %0.3f \n", a1, a2)
}
