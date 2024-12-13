package main

func mapInt(arr []int, f func(int) int) []int {
	result := make([]int, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}
func mapAny[K, V any](arr []K, f func(K) V) []V {
	result := make([]V, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// Hàm Sum sử dụng interface Number.
func Sum[T Number](numbers []T) T {
	var sum T
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func main() {
	// arrayTest := []int{1, 2, 3, 4, 5, 6}
	// rs := mapInt(arrayTest, func(v int) int {
	// 	return v * 2
	// })
	// rs2 := mapAny(arrayTest, func(v int) int {
	// 	return v * 2
	// })
	// fmt.Println(rs2)
	// fmt.Println(rs)
	ints := []int{1, 2, 3, 4, 5}
	floats := []float64{1.1, 2.2, 3.3}

	// Gọi hàm Sum với int
	println(Sum(ints))

	// Gọi hàm Sum với float64
	println(Sum(floats))
}
