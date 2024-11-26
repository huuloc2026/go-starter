package main

import (
	"fmt"
)

// const a = 42

// const (
// 	RED    = 0
// 	YELLOW = 1
// 	BLUE   = 2
// )

// iota
// const (
// 	RED = iota
// 	YELLOW
// 	BLUE
// ) // block scope

const (
	_ = iota // can apply operators ex: iota + 5, 1 << iota
	RED
	YELLOW
	BLUE
) // block scope

func main() {
	// const i = "Hello"
	// const i = math.Sqrt(4) -> error
	// const a = "Shadow"
	// fmt.Printf("%v - type: %T", a, a)
	// fmt.Printf("%v - type: %T \n", i, i)

	//iota
	fmt.Printf("%v - type: %T \n", RED, RED)
	fmt.Printf("%v - type: %T \n", YELLOW, YELLOW)
	fmt.Printf("%v - type: %T \n", BLUE, BLUE)
}
