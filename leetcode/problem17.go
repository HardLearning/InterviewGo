package main

/**
* @author moorewu
* @E-mail moorewu@tencent.com
* @date 2020/2/26 4:00 下午
 */

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	mp := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	result := mp[digits[0]]
	for i := 1; i < len(digits); i++ {
		temp := make([]string, 0)
		for j := 0; j < len(result); j++ {
			for k := 0; k < len(mp[digits[i]]); k++ {
				temp = append(temp, result[j]+mp[digits[i]][k])
			}
		}
		result = temp
	}

	return result
}

//func main() {
//	fmt.Println(letterCombinations("23"))
//}
