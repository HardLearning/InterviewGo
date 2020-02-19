package main

import (
	"math"
)

func reverse(x int) int {
	var num int
	for x != 0 {
		num = num * 10 + x % 10
		x /= 10
	}
	if num > math.MaxInt32 || num < math.MinInt32 {
		num = 0
	}
	return num
}
