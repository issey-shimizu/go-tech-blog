package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	ans := -1
	for i, _ := range s{
		if s[i] == 'a' {
			ans = i + 1
		}
	}
	fmt.Println(ans)
}


