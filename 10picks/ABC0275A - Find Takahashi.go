package main

import "fmt"

func main() {
	var n, h int 
	fmt.Scan(&n)
	id := 0
	max := 0

	for i := 0; i < n; i++{
		fmt.Scan(&h)
		if h > max {
			max = h
			id = i + 1
		}
	}
 	fmt.Println(id)
}