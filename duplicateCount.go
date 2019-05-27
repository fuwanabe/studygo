package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(duplicateCount("abcdeaB11"))
}

func duplicateCount(s1 string) (c int) {
	m := make(map[rune]int)
	for _, v := range strings.ToLower(s1) {
		m[v]++
		if m[v] == 2 {
			c++
		}
	}
	return
	//strMap := make(map[string]int)
	//strArray := strings.Split(s1, "")
	//for _, char := range strArray {
	//	downChar := strings.ToLower(char)
	//	if _, ok := strMap[downChar]; ok {
	//		strMap[downChar]++
	//	} else {
	//		strMap[downChar] = 1
	//	}
	//}
	//result := 0
	//for _, v := range strMap {
	//	if v > 1 {
	//		result++
	//	}
	//}
	//return result
}
