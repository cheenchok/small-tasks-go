package main

import (
	"fmt"
)

func main() {
	var n, x1, x2, x3 int
	var count int = 0
	fmt.Scan(&n)

	for n := 100; n <= 999; n++ {
		x1 = n / 100
		x2 = (n % 100) / 10
		x3 = n % 10
		if x1+x2+x3 == 8 {
			count++
		}
	}
	fmt.Println(count)
}
