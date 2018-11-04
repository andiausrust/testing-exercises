package main

import "fmt"

func main() {

	fmt.Println("2+3 = ", mySumTable(2, 3))
	fmt.Println("3+3 = ", mySumTable(3, 3))
}

func mySumTable(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
