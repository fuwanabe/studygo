package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 1, 1, 1, 2, 3, 6, 7, 8, 9}
	fmt.Println(findClosestElements(nums, 7, 4))
}

func findClosestElements(arr []int, k int, x int) []int {
	//len(arr)-kの-kの意味はarr[i+k] で配列のインデックスを超えないようにするため
	i := sort.Search(len(arr)-k, func(i int) bool {
		fmt.Printf("i : %v, x-arr[i] : %v, arr[i+k]-x : %v \n", i, x-arr[i], arr[i+k]-x)
		return x-arr[i] <= arr[i+k]-x
	})
	fmt.Println(i)
	//iからk個の範囲まで取得
	return arr[i : i+k]
	//まずｘを見つける
	//xIndex := sort.Search(len(arr), func(i int) bool {
	//	return arr[i] >= x
	//})
	//fmt.Println(xIndex)
	//needsNumber := k - 1
	//res := make([]int, 0)
	//res = append(res, arr[xIndex])
	//sideIndex := 1
	//for needsNumber > 0 {
	//	if xIndex-sideIndex >= 0 {
	//		res = append(res, arr[xIndex-sideIndex])
	//		needsNumber--
	//		if needsNumber == 0 {
	//			break
	//		}
	//	}
	//	if xIndex+sideIndex <= len(arr) {
	//		res = append(res, arr[xIndex+sideIndex])
	//		needsNumber--
	//	}
	//	sideIndex++
	//}
	//sort.Slice(res, func(i, j int) bool {
	//	return res[i] < res[j]
	//})
	//
	//return res
}
