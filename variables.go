package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int
	x = 5
	y := 3
	fmt.Println(x, y)
	myvalue, err := strconv.ParseInt("9", 0, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myvalue)
	}
	m := make(map[string]float64)
	m["edad"] = 7.5
	fmt.Println(m)
	s := []int{5, 3, 4, 6}
	for _, value := range s {
		fmt.Println(value)
	}
	s = append(s, 5)
	fmt.Print(s)
}
