package main

import (
	"algorithm/sorting"
	"fmt"
)

func main() {
	array := []int{
		34, 59, 61, 98, 2,
		45, 3, 7, 11, 18,
		56, 87, 21, 9, 35,
		78, 30, 20, 53, 49,
		19, 5, 22, 38, 1,
	}

	fmt.Println(array)

	// Bubble sort
	result := sorting.BubbleSort(array)
	fmt.Println(result)
}
