package main

import "fmt"

func main() {
	half(1)
	half(2)
}

func half(num int) {
	n := num / 2

	if n%2 == 0 {
		fmt.Println(n, true)
	} else {
		fmt.Println(n, false)
	}

}
