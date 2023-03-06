package main

import (
	"fmt"
	"time"
)

func main() {
	x := 5
	y := func() int {
		return x + 2
	}()
	fmt.Println(y)
	c := make(chan int)
	go func() {
		fmt.Println("Starting function")
		time.Sleep(5 * time.Second)
		fmt.Println("ending function")
		c <- 1
	}()
	<-c
}
