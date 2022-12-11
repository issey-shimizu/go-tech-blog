package main

import (
	"fmt"
	"strings"
)

func main() {
	var n string
	fmt.Scan(&n)
	fmt.Println(strings.Count(n, "1"))
}