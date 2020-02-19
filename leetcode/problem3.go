package main

import "math"

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]byte)
	sLen := len(s)
	start,end := 0,0
	repeatCount := 0

	for start < sLen && end < sLen {
		tmpString := s[end]
		if _, ok := m[tmpString]; !ok {
			m[s[end]] = s[end]
			end++
			repeatCount = int(math.Max(float64(repeatCount), float64(end - start)))
		} else {
			delete(m, s[start])
			start++
		}
	}
	return repeatCount
}
