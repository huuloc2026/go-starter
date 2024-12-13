package main

import "fmt"

func main() {
	s := []int{2, 3, 4, 5, 6, 7, 8}
	fmt.Print(s)

	// pointStudent1 := 7
	// pointStudent2 := 10
	// fmt.Printf("%v - %T \n", pointStudent1, pointStudent1)
	// fmt.Printf("%v - %T", pointStudent2, pointStudent2)
	// points := [3]int{7, 5, 10}
	// var points [3]int // [3]string
	// var points [3]int // [3]string
	// fmt.Printf("%v - %T", points, points) //-> [0 0 0] [3] int
	// points[0] = 10
	// points[1] = 11
	// points[2] = 21
	// points[0] = "Jake"
	// points[1] = "Jake"
	// points[2] = "Jake"
	// fmt.Printf("%v - %T", points, points) //-> [0 0 0] [3] int
	// fmt.Println(len(points)) //-> [0 0 0] [3] int
	// points := [...]int{7, 5, 10} // auto init length of array
	// fmt.Printf("%v - %T", points, points)
	// :::NEW ARRAY:::
	// a := [3]int{1, 2, 3}
	// // b := a  // copy
	// b := &a // pointer
	// b[0] = 999
	// fmt.Println(a)
	// fmt.Println(b)
	// /// **** ::::SLICES SMILILAR ARRAY BUT DYNAMIC SIZE
	// a := []int{2, 5, 10}
	// b := a
	// b[0] = 999
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println("len::", len(a))
	// fmt.Println("cap::", cap(a))

	// a := []int{2, 5, 10, 12, 45, 50}
	// b := a[:]
	// c := a[3:]
	// d := a[:5]
	// e := a[3:5]
	// e[1] = 100
	// fmt.Printf("a:::%v - %v - %v \n ", a, len(a), cap(a))
	// fmt.Printf("b:::%v - %v - %v \n ", b, len(b), cap(b))
	// fmt.Printf("c:::%v - %v - %v \n ", c, len(c), cap(c))
	// fmt.Printf("d:::%v - %v - %v \n ", d, len(d), cap(d))
	// fmt.Printf("e:::%v - %v - %v \n ", e, len(e), cap(e))
	// a := make([]int, 10, 20) // 10 -> length, 20 -> cap
	// a := make([]int, 0)
	// fmt.Printf("a:::%v - %v - %v \n ", a, len(a), cap(a))
	// a = append(a, 1)                 // increase length&cap
	// a = append(a, 5, 6, 7, 8, 9, 10) // increase length&cap
	// spread := []int{99, 98, 99, 99}
	// // a = append(a, []int{99, 98, 99, 99}...) // increase length&cap
	// a = append(a, spread...) // increase length&cap
	// // a = append(a, 10)      // increase length&cap
	// fmt.Printf("a:::%v - %v - %v \n ", a, len(a), cap(a))

	// a := []int{1, 2, 3, 4, 5, 6}
	// // b := a[:5] // pop
	// // b := a[1:] // shift
	// b := a[1:4]
	// // b := append(a[:2], a[3:]...)
	// fmt.Println(b)
}
