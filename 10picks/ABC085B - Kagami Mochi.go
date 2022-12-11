package main

import "fmt"

func main() {
	var n, d int
	fmt.Scanf("%d", &n)
	inputMap := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &d)
		inputMap[d] = 1
	}
	fmt.Println(len(inputMap))
}
