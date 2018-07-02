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

/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
Solution == 233168
*/
