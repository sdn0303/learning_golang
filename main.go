package main

import (
	"fmt"
	"github.com/sdn0303/learning/sort_algorithm"
)

func main() {
	numbers := []int{9, 2, 1, 5, 7, 3, 8, 4, 6, 0}

	result := sort_algorithm.BubbleSort(numbers)
	fmt.Printf("Bubble Sort: %v \n", result)
	fmt.Println("------------------------")

	result = sort_algorithm.SelectionSort(numbers)
	fmt.Printf("Selection Sort: %v \n", result)
	fmt.Println("------------------------")
}
