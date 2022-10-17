package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	maxValue := array[0]
	for i := 0; i < 5; i++ {
		if array[i] > maxValue {
			maxValue = array[i]
		} else if array[i] == 0 && array[i] > maxValue {
			maxValue = array[i]
		}
	}
	fmt.Print(maxValue)
}
