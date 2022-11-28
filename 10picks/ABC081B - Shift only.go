package main

import (
	"fmt"
)

func main() {
	var 
	var num, a, b, c int

	fmt.Scanf("%d", &num)
	fmt.Scanf("%d %d %d", &a, &b, &c)
	ary := [*&num]int{*&a,*&b,*&c}
	fmt.Println(ary)

}