package order

import "fmt"

//go run order/bubblesort.go
//example:  https://www.youtube.com/watch?v=8RkZoBZNNgI

//bubble sort is a order algorithm that compares two elements and swap them if the first one is bigger than the second one

// BubbleSort is a function that order a slice of integers
func BubbleSort(input []int) {
	fmt.Println("Unsorted:", input)
	var changed bool
	for i := 0; i < len(input)-1; i++ {
		changed = false
		for j := 0; j < len(input)-1; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
				changed = true
			}
		}
		if !changed {
			break
		}
	}
	fmt.Println("Sorted:  ", input)
}
