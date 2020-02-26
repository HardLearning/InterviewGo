package main

/**
* @author moorewu
* @E-mail moorewu@tencent.com
* @date 2020/2/26 3:50 下午
 */

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}
	result := strs[0]

	for i := 1; i < length; i++ {
		mn := len(result)
		if len(strs[i]) < mn {
			mn = len(strs[i])
		}

		tmp := ""
		for j := 0; j < mn; j++ {
			if result[j] == strs[i][j] {
				tmp += string(result[j])
			} else {
				break
			}
		}

		result = tmp
	}

	return result
}

//func main() {
//	strs := []string{"flower", "flow", "flight"}
//	fmt.Println(longestCommonPrefix(strs))
//}
