package main

import (
	"fmt"
)

func main() {
	var N, A, B int
	var result int = 0
	fmt.Scanf("%d %d %d", &N, &A, &B)

	for i := 1; i <= N; i++ {
		if A <= calcSum(i) && calcSum(i) <= B {
			result += i
		}
	}

	fmt.Println(result)
}

func calcSum(x int) int {
	var sum int = 0
	for x > 0 {
		sum += x % 10
		x = x / 10
	}
	return sum
}
