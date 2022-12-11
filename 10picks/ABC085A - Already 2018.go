package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println("2018" + str[4:])
}
