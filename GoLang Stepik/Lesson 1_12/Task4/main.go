package main

import "fmt"

func main() {
	var value, n int
	fmt.Scan(&n)
	arr := make([]int, n)
	if n >= 1 && n <= 100 {
		for i := 0; i < n; i++ {
			fmt.Scan(&value)
			arr[i] = value
		}
		for i, item := range arr {
			if i%2 == 0 {
				fmt.Print(item, " ")
			} else if i == 0 {
				fmt.Print(item, " ")
			}
		}
	}

}
