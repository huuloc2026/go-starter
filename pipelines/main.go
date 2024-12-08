package main

import (
	"fmt"
)

// Worker 1: Tạo số từ 1 đến n
func generateNumbers(n int, out chan int) {
	for i := 1; i <= n; i++ {
		out <- i
	}
	close(out)
}

// Worker 2: Tính bình phương của các số
func squareNumbers(in chan int, out chan int) {
	for num := range in {
		out <- num * num
	}
	close(out)
}

func main() {
	numbers := make(chan int)
	squares := make(chan int)

	go generateNumbers(5, numbers)
	go squareNumbers(numbers, squares)

	// Nhận và in kết quả
	for square := range squares {
		fmt.Println(square)
	}
}
