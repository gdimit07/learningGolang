package main

import "fmt"

func printMap(m map[string]int) {
	for spelled, num := range m {
		fmt.Println("spelled number:", spelled, "=>numerical value:", num)
	}
}

func main() {

	//below does not work, map should be initialized  with make or colors := map[<type>]<type>
	//var colors_a map[string]string

	colors_b := make(map[int]int)

	colors_b[1] = 2
	colors_b[3] = 4
	delete(colors_b, 3)

	fmt.Println(colors_b)

	colors := map[string]int{
		"one": 1,
		"two": 2,
	}

	colors["three"] = 3

	//fmt.Println(colors)
	printMap(colors)
}
