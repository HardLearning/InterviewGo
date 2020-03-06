package main

import (
	"fmt"
	"math"
)

/**
* @author moorewu
* @E-mail moorewu@tencent.com
* @date 2020/3/5 11:53 下午
 */

func judgeSquareSum(c int) bool {
	i, j := 0, int(math.Sqrt(float64(c)))+1
	for i <= j {
		switch {
		case i*i+j*j > c:
			j--
		case i*i+j*j == c:
			return true
		default:
			i++
		}
	}
	return false
}

func main() {
	fmt.Println(judgeSquareSum(5))
}
