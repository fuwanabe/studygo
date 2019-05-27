package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 1, 1, 1, 2, 3, 6, 7, 8, 9}
	fmt.Println(searchRange(nums, 2))
}

func searchRange(nums []int, target int) []int {
	//右端を探す
	start := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	//左橋（ターゲットの1つ以上多いの最初の数値）
	end := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target+1
	})
	return []int{start, end - 1}
}
