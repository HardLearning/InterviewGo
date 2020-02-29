package main

/**
* @author moorewu
* @E-mail moorewu@tencent.com
* @date 2020/2/29 3:12 下午
 */

//问题描述
//请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。
//给定一个string，请返回一个string，为翻转后的字符串。保证字符串的长度小于等于5000。

func flipString(s string) (string, bool) {
	str := []rune(s)
	length := len(str)
	if length > 5000 {
		return string(str), false
	}
	for i := 0; i < length/2; i++ {
		str[i], str[length-1-i] = str[length-1-i], str[i]
	}
	return string(str), true
}
