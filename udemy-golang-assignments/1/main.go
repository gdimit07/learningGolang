package main

import "fmt"

func main() {

	//fmt.Println("Assignment 1 - Even and Odd")

	numbers := []int{}
	var i int

	//create slice
	for i = 0; i <= 10; i++ {

		numbers = append(numbers, i)
	}

	//iterate through slice and check its elements
	for _, element := range numbers {

		if element%2 == 0 {
			fmt.Println(element, "is even")
		} else {
			fmt.Println(element, "is odd")
		}
	}
}
