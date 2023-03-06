package main

import "fmt"

func main() {
	fmt.Println("help")
	// c := make(chan int)
	// go sleep(c)
	// <-c
	g := 52
	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}

func sleep(c chan int) {
	fmt.Print("zzzz tengo sueño")
	c <- 1
}
