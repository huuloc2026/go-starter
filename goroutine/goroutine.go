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
// 		fmt.Println(name, i)
// 		time.Sleep(time.Second)
// 	}
// }
