package main

import "fmt"

func main() {
	// println(swap(5, 15))
	// structlearn()
	pointerLearn()
}

func swap(x, y int) (int, int) {
	return y, x
}
func structlearn() {
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Huu Loc", Age: 5}
	fmt.Println(p)
}
func pointerLearn() {
	var a *int
	x := 15
	a = &x
	fmt.Println(*a)
}
