package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
  	fmt.Scanf("%s", &s)
	  slice := strings.Split(s, "")
	  count := 0
  	for i := 0; i < len(slice); i++ {
    if slice[i] == "1" {
			count++
		}
  }
	fmt.Printf("%d",count)
}
