package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Points([]string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"}))
}

func Points(games []string) int {
	total := 0
	for _, game := range games {
		x, _ := strconv.ParseInt(game[:1], 10, 0)
		y, _ := strconv.ParseInt(game[2:], 10, 0)
		switch {
		case x > y:
			total += 3
		case x == y:
			total += 1
		case x < y:
		}
	}
	return total
}
