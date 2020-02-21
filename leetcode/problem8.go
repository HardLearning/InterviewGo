package main

import (
	"math"
	"strings"
)

func isDigit(x uint8) bool {
	return x >= '0' && x <= '9'
}

func isSymbol(x uint8) bool {
	return x == '+' || x == '-'
}

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}

	if !isSymbol(str[0]) && !isDigit(str[0]) {
		return 0
	}

	symbol := 1
	if isSymbol(str[0]) {
		if str[0] == '-' {
			symbol = -1
		}
		str = str[1:]
	}

	result := 0

	for i:=0; i < len(str); i++ {
		if !isDigit(str[i]) {
			break
		}
		result = result * 10 + int(str[i] - '0')

		if symbol * result > math.MaxInt32 {
			return math.MaxInt32
		}

		if symbol * result < math.MinInt32 {
			return math.MinInt32
		}
	}

	result = result * symbol
	return result
}

//func main() {
//	str1 := "-100"
//	fmt.Println(myAtoi(str1))
//
//	str2 := "100"
//	fmt.Println(myAtoi(str2))
//
//	str3 := "hello 100"
//	fmt.Println(myAtoi(str3))
//
//	str4 := "-91283472332"
//	fmt.Println(myAtoi(str4))
//}
