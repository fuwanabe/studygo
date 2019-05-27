package main

import "fmt"

func main() {
	nums := []int{0, 1, 1, 1, 2, 3, 6, 7, 8, 9}
	fmt.Println(findMin(nums))
}

func findMin(nums []int) int {
	min := nums[0]
	for i, v := range nums {
		if 0 < i {
			if min > v {
				min = v
			}
		}
	}
	return min
}
