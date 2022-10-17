package main

import "fmt"

func main() {
	var n, num int
	fmt.Println("Ввод n:")
	fmt.Scan(&n)
	arr := make([]int, n)
	if n >= 4 {
		for i := 0; i < n; i++ {
			fmt.Scan(&num)
			arr[i] = num
		}
		fmt.Print(arr[3])
	}
}
