package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v \n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}

// package goroutine

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func goroutine() {
// 	var wg = sync.WaitGroup{}
// 	wg.Add(2)
// 	go func() {
// 		count("Sheep")n k 0, 0,l0
// 		wg.Done()
// 	}()
// 	go func() {
// 		count("Elephant")
// 		wg.Done()
// 	}()
// 	wg.Wait()
// 	fmt.Println("Done")
// }
// func count(name string) {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println(name, i)s
// 		time.Sleep(time.Second)
// 	}
// }
