package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	fmt.Println("Hello goroutine", i)
}

func showNumber(i int) {
	fmt.Println(i)
}

func main() {
	ch := make(chan int)
	go func() {
		ch <- 10
	}()
	var x int
	go func() { x = <-ch }()
	time.Sleep(time.Millisecond)
	fmt.Println(x)

	close(ch)
	// go func() {
	// 	defer fmt.Println("A.defer")
	// 	func() {
	// 		defer fmt.Println("B.defer")
	// 		runtime.Goexit()
	// 		defer fmt.Println("C.defer")
	// 		fmt.Println("B")
	// 	}()
	// 	fmt.Println("A")
	// }()

	// for i := 0; i < 10; i++ {
	// 	go showNumber(i)
	// }

	// runtime.Gosched()
	// fmt.Println("Haha")

	// go func(s string) {
	// 	for i := 0; i < 2; i++ {
	// 		fmt.Println(s, i)
	// 	}
	// }("world")

	// for i := 0; i < 2; i++ {
	// 	runtime.Gosched()
	// 	fmt.Println("Hello", i)
	// }

	// go func() {
	// 	i := 0
	// 	for {
	// 		i++
	// 		fmt.Printf("new goroutine: i = %d\n", i)
	// 		time.Sleep(time.Second)
	// 		if i == 2 {
	// 			break
	// 		}
	// 	}
	// }()
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go hello(i)
	// }
	// fmt.Println("第一")
	// wg.Wait()
	// fmt.Println("main goroutine done")
	// time.Sleep(time.Second)
}
