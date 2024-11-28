package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	CountEven(array)

}
func CountEven(array []int) {
	var countEven int
	for _, value := range array {
		if value%2 == 0 {
			countEven++
		}
	}
	fmt.Println("count::", countEven)
	fmt.Println("countOdd::", len(array)-countEven)
}
