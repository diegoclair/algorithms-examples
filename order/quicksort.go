package order

import "fmt"

//go run order/quicksort.go
//quicksort is a order algorithm that uses a pivot to divide the slice in two parts, one with elements smaller than the pivot and another with elements bigger than the pivot

// example: https://www.youtube.com/watch?v=wU7Q8Z51MUI

func QuickSort(input []int) {
	fmt.Println("Unsorted: ", input)
	quicksort(input, 0, len(input)-1)
	fmt.Println("Sorted:  ", input)
}

func quicksort(input []int, left int, right int) {
	if left < right {
		pivot := partition(input, left, right)
		quicksort(input, left, pivot-1)
		quicksort(input, pivot+1, right)
	}
}

func partition(input []int, left int, right int) int {
	pivot := input[right]
	i := left - 1

	for j := left; j < right; j++ {
		if input[j] <= pivot {
			i++
			input[i], input[j] = input[j], input[i]
		}
	}

	input[i+1], input[right] = input[right], input[i+1]

	return i + 1
}
