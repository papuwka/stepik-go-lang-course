package main

import "fmt"

func main() {
	var value, n, k int
	fmt.Scan(&n)
	arr := make([]int, n)
	if n >= 1 && n <= 100 {
		for i := 0; i < n; i++ {
			fmt.Scan(&value)
			arr[i] = value
		}
		for _, item := range arr {
			if item > 0 {
				k++
			}
		}
		fmt.Println(k)
	}

}
