package main

import "fmt"

func main() {
	fmt.Println(mySqrt(8))
}

func mySqrt(x int) int {
	xf := float64(x)
	number := float64(x)
	before := float64(0)
	for number != before {
		before = number
		number = (number + xf/number) / 2.0
	}
	return int(number)
}
