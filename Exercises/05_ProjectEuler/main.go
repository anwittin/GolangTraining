package main

import "fmt"

func main() {
	var multiples int
	for i := range [1000]int{} {
		if i%3 == 0 || i%5 == 0 {
			multiples += i
		}
	}
	fmt.Println(multiples)
}
