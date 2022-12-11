package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var t, x, y int
		fmt.Scan(&t, &x, &y)
		if t < x+y || t%2 != (x+y)%2 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
