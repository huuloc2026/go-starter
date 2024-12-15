package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Second * 2)
		ch <- 1
	}()
	fmt.Println(<-ch)
}
