package main

import (
	"fmt"
)

func main() {
	var x1, x2 int
	var separator = ""
	for n := 10; n <= 99; n++ {
		x1 = n / 10
		x2 = n % 10
		if n == x1*x2*2 {
			fmt.Print(separator)
			fmt.Print(n)
			separator = ","
		}
	}
}
