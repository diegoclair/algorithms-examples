package main

import "fmt"

//go run order/quicksort.go
//quicksort is a order algorithm that uses a pivot to divide the slice in two parts, one with elements smaller than the pivot and another with elements bigger than the pivot

// example: https://www.youtube.com/watch?v=wU7Q8Z51MUI
var input = []int{48, 9, 86, 32, 68, 57} //, 82, 63, 70, 37, 34, 83, 27, 19, 97, 96, 17}

func main() {
	fmt.Println("Unsorted:", input)
	QuickSort(input)
	fmt.Println("Sorted:", input)
}

func QuickSort(input []int) {
	quicksort(input, 0, len(input)-1)
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
