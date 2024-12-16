package main

import (
	"fmt"
	"slices"
)

func main() {
	m := 4
	n := 3
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, m, nums2, n)
}
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	k := (m + n) - 1
	j := n - 1
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
	slices.Sort(nums1)
	fmt.Println(nums1)
	return nums1
}
