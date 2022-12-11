package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanf("%d %d\n", &a, &b)
	if ((a * b) % 2) == 0 {
		fmt.Println("Even")
		return
	}
	fmt.Println("Odd")
}
