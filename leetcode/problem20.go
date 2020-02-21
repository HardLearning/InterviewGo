package main

import "fmt"

// 数据结构简单题， 模拟栈操作
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	str := ""
	for i:=0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			str += string(s[i])
			continue
		}
		if s[i] == ')' {
			if len(str) == 0 || str[len(str)-1:] != "(" {
				return false
			} else {
				str = str[:len(str)-1]
			}
		}
		if s[i] == ']' {
			if len(str) == 0 || str[len(str)-1:] != "[" {
				return false
			} else {
				str = str[:len(str)-1]
			}
		}
		if s[i] == '}' {
			if len(str) == 0 || str[len(str)-1:] != "{" {
				return false
			} else {
				str = str[:len(str)-1]
			}
		}
	}
	return len(str) == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("("))
}
