package main

import "fmt"

func divide(datas []int) []int {
	for i, data := range datas {
		if data%2 != 0 {
			datas[i] = 0
		} else {
			datas[i] = data / 2
		}
	}
	return datas
}

func main() {
	var num int
	fmt.Scan(&num)

	datas := make([]int, 0, num)
	for i := 0; i < num; i++ {
		var d int
		fmt.Scan(&d)
		datas = append(datas, d)
	}

	count := 0
	for {
		datas = divide(datas)
		if datas == nil {
			break
		}
		count++
	}
	fmt.Println(count)
}
