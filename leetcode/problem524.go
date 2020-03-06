package main

import (
	"fmt"
	"strings"
)

/**
* @author moorewu
* @E-mail moorewu@tencent.com
* @date 2020/3/6 10:01 上午
 */

func findLongestWord(s string, d []string) string {
	answer := ""
	for _, ss := range d {
		if len(ss) < len(answer) {
			continue
		}
		if len(ss) == len(answer) && strings.Compare(ss, answer) == 1 {
			continue
		}
		if SubStr(s, ss) {
			answer = ss
		}
	}
	return answer
}

func SubStr(s, target string) bool {
	l1, l2 := len(s), len(target)
	i, j := 0, 0
	for i < l1 && j < l2 {
		if s[i] == target[j] {
			i++
			j++
		} else {
			i++
		}
	}
	return j >= l2
}

func main() {
	d := []string{"ale", "apple", "monkey", "plea"}
	s := "abpcplea"
	ret := findLongestWord(s, d)
	fmt.Println(ret)
}
