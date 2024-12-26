package main

func main() {
	numbers := []int{2, 3, 4}
	target := 4
	twoSum(numbers, target)
}

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	// If no solution exists
	return []int{-1, -1}
}
