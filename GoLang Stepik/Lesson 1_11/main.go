package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scan(&a)
	if a <= 0 {
		fmt.Printf("число %1.2f не подходит", a)
	} else if a > 10000 {
		fmt.Printf("%e", a)
	} else if a > 0 && a < 10000 {
		b = a * a
		fmt.Printf("%2.4f", b)
	}
}
