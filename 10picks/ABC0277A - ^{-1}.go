package main

import "fmt"

func main() {
	var N, X, P int
	fmt.Scan(&N, &X)

	for i:=1 ; i <= N ; i++{
		fmt.Scan(&P)
		if P == X {
			fmt.Println(i)
		}
	}
}
