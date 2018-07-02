package main

import "fmt"

func main() {
	n := max(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func max(x ...int) int {

	var total int
	for _, v := range x {
		if v > total {
			total = v
		}
	}
	return total
}
