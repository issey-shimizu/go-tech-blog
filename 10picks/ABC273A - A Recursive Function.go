package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	y := 1
	for i := 2 ; i <= n ;i++ {
		y = y * i
	}
	fmt.Println(y)
}





