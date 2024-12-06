package main

import "fmt"

type Car struct {
	Name   string
	Speed  int
	Wheels int
}
type Bike struct {
	Name   string
	Speed  int
	Wheels int
}

func (c Car) Run() string {
	return c.Name + " chạy bằng 4 bánh!"
}

func (b Bike) Run() string {
	return b.Name + " chạy bằng 2 bánh!"
}

type Vehicle interface {
	Run() string
}

func main() {
	var v Vehicle
	car := Car{Name: "Toyota", Speed: 120, Wheels: 4}
	bike := Bike{Name: "Yamaha", Speed: 80, Wheels: 2}

	v = car
	fmt.Println(v.Run()) // Kết quả: Toyota chạy bằng 4 bánh!

	v = bike
	fmt.Println(v.Run()) // Kết quả: Yamaha chạy bằng 2 bánh!
}
