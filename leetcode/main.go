package main

import "fmt"

func main() {
	array1 := []int{1, 1, 1, 2, 2, 3, 4, 5, 6}
	removeDuplicate(array1)
}

func removeDuplicate(nums []int) []int {
	// [1, 1, 1, 2, 2, 3, 4, 5, 6]
	//  L
	//          R                  // [1,2]
	//			L
	//          R

	left := 0
	right := 1
	result := []int{nums[left]}
	for right < len(nums) {
		if nums[right] > nums[left] {
			result = append(result, nums[right])
			left = right
		}
		right++
	}
	fmt.Println("result:::", result)
	return result
}
