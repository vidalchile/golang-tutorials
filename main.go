package main

import "fmt"

func main() {
	var ages1 [3]int = [3]int{20, 25, 30}
	var ages2 = [3]int{20, 25, 30}

	fmt.Println(ages1, ages2)

	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	names[1] = "luigi"

	fmt.Println(names)

	fmt.Println(ages2, len(ages2))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:4] // doesn't include pos 4 element
	rangeTwo := names[2:]  //includes the last element
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)
	fmt.Printf("the type of rangeOne is %T \n", rangeOne)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)
}
