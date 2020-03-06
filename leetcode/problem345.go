package main

/**
* @author moorewu
* @E-mail moorewu@tencent.com
* @date 2020/3/6 12:15 上午
 */

func reverseVowels(s string) string {
	t := []byte(s)
	m := make(map[byte]bool)
	m['a'], m['e'], m['i'], m['o'], m['u'] = true, true, true, true, true
	m['A'], m['E'], m['I'], m['O'], m['U'] = true, true, true, true, true
	for i, j := 0, len(t)-1; i < j; i, j = i+1, j-1 {
		for i < j && !m[t[i]] {
			i++
		}
		for i < j && !m[t[j]] {
			j--
		}
		if i >= j {
			break
		}
		t[i], t[j] = t[j], t[i]
	}
	return string(t)
}
