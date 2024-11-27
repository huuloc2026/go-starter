package main

import "fmt"

func main() {
	//var a int8 = 10 //1010 -> 2^3 -> 0000 0100
	// var b float64
	// b = 13.3e72 // 13.3*10^72

	//fmt.Printf("%v,%T\n", a<<3, a<<3)

	// var a float32 = 13.25
	// var b float32 = 2.15
	// fmt.Printf("%v,%T\n", a+b, a+b)
	// fmt.Printf("%v,%T\n", a-b, a-b)
	// fmt.Printf("%v,%T\n", a*b, a*b)
	// fmt.Printf("%v,%T\n", a/b, a/b)
	var a complex128 = complex(5, 12)
	fmt.Printf("%v,%T \n", a, a)
	fmt.Printf("%v,%T \n", real(a), real(a))
	fmt.Printf("%v,%T", imag(a), imag(a))

}
