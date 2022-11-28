package main

import (
	"fmt"
)

func main() {
    var num, a, b, c int
	fmt.Scanf("%d", &num)
	fmt.Scanf("%d %d %d", &a, &b, &c)
	ary := [...]int{*&a,*&b,*&c}
    num := len(ary)
    count := 0
  	for i := 0; i < num; i++ {
      sum := a + b + c
      split := sum / 2
      if spilt % 2 == 0 {
			count++
	   }
		}
  }
  fmt.Printf("%d",count)

if sss == 0{

xq
}

}