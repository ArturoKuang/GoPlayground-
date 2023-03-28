package main

import (
	"fmt"
)

func main() {
	msgch := make(chan int)

	// go func() {
	// 	fmt.Println("0")
	// 	msg := <-msgch
	// 	fmt.Println(msg)
	// }()

	// go func() {
	// 	fmt.Println("1")
	// 	msgch <- 10
	// 	msgch <- 20
	// 	fmt.Println("2")
	// 	msgch <- 30
	// }()

	// msg := <-msgch
	// fmt.Println(msg)

	msg := <-msgch

	// time.Sleep(20 * time.Millisecond)
	fmt.Println("Hello, world!")
}
