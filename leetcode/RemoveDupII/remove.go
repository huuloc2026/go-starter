package main

import "fmt"

func main() {
	array1 := []int{0, 0, 0, 0, 1, 2, 2, 2, 3}

	fmt.Println(removeDuplicate(array1))
}

func removeDuplicate(nums []int) int {
	k := 2
	i := 0
	for _, num := range nums {

		if i < k || num != nums[i-k] {
			nums[i] = num
			i++
		}
	}
	return i
	// Quá trình thực thi từng vòng lặp
	// Giả sử nums = [1, 1, 1, 2, 2, 3]:

	// Vòng 1:
	// num = 1, i = 0, i < k → Ghi 1 vào nums[0], tăng i = 1.
	// nums {0,1,}
	// Vòng 2:
	// num = 1, i = 1, i < k → Ghi 1 vào nums[1], tăng i = 2.

	// Vòng 3:
	// num = 1, i = 2, num == nums[i-k] → Bỏ qua (không ghi).

	// Vòng 4:
	// num = 2, i = 2, num != nums[i-k] → Ghi 2 vào nums[2], tăng i = 3.

	// Vòng 5:
	// num = 2, i = 3, num != nums[i-k] → Ghi 2 vào nums[3], tăng i = 4.

	// Vòng 6:
	// num = 3, i = 4, num != nums[i-k] → Ghi 3 vào nums[4], tăng i = 5.
}
