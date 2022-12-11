package main

import (
	"fmt"
)

func main() {
	var n, a int
	fmt.Scan(&a)
	sum := 0

	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		sum += a
	}
	fmt.Println(sum)
}
