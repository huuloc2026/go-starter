package main

import (
	"fmt"
)

func main() {
	array := []int{2, 3, 4, 6, 1, 8, 9}
	sortedArray := quicksort(array)
	fmt.Println(sortedArray)
}
func quicksort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	pivot := array[0]
	less := []int{}
	greater := []int{}
	for _, v := range array[1:] {
		if v < pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}
	return append(append(quicksort(less), pivot), quicksort(greater)...)
}
