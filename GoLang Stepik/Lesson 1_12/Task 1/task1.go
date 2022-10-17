package main

import "fmt"

func main() {
	var workArray [10]uint8
	var num, indexFirst, indexSecond uint8
	for i := 0; i < len(workArray); i++ {
		fmt.Scan(&num)
		workArray[i] = num
	}
	for i := 0; i < 3; i++ {
		fmt.Scan(&indexFirst, &indexSecond)
		workArray[indexFirst], workArray[indexSecond] = workArray[indexSecond], workArray[indexFirst]
	}
	for i := 0; i < len(workArray); i++ {
		fmt.Print(workArray[i], " ")
	}
}
