package main

import "fmt"

func main() {
	v := 5
	p := &v
	fmt.Println(v)
	fmt.Println(&v)
	fmt.Println(&p)
}
