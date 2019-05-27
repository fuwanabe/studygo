package main

import "fmt"

func main() {
	nums := []int{0, 1, 1, 1, 2, 3, 6, 7, 8, 9}
	fmt.Println(findPeakElement(nums))
}

func findPeakElement(nums []int) int {
	result := 0
	max := len(nums) - 1
	for i, v := range nums {
		if i > 0 { //0を除外
			if i < max && nums[i-1] < v && v > nums[i+1] { //両隣より大きい
				return i
			} else if i == max && nums[i-1] < v { //一番右が左どなりより大きい
				result = i
			}
		}
		fmt.Printf("%v回目\n", i)
	}
	return result

}
