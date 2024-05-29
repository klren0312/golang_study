package main

import (
	"fmt"
	"log"
	"time"

	"zzes1314.cn/greeting"
)

func trace(s string) string {
	fmt.Println("entering: ", s)
	return s
}

func un(s string) {
	fmt.Println("leaving: ", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func deferTest() {
	defer fmt.Println("close")
	fmt.Println("do 1")
	return
}

func Announce(message string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println(message)
	}()
}

func main() {
	log.SetPrefix("greeting: ")
	log.SetFlags(0)
	// fmt.Println(quote.Go())
	message, err := greeting.Hello("test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"a1", "b1", "c1"}
	messages, err1 := greeting.Hellos(names)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(messages)
	// b()
	// deferTest()
	// var (
	// 	firstName string = "job"
	// 	lastName  string = "lala"
	// 	age       int
	// )
	// fmt.Println("Hello World!")
	// fmt.Println(firstName, lastName, age)
	// fmt.Println(math.MaxFloat32, math.MaxFloat64)
	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println("sum is", sum)

	// for pos, char := range "日本x80	語" {
	// 	fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	// }

	// for i := 0; i < 5; i++ {
	// 	defer fmt.Printf("%d", i)
	// }
}
