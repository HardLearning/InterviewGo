package main

func romanToInt(s string) int {
	specialNumberMap := map[string]int{"IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}
	normalNumberMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	result := 0

	for len(s) > 0 {
		if len(s) > 1 {
			str := s[0:2]
			if val, ok := specialNumberMap[str]; ok {
				result += val
				s = s[2:]
			} else {
				result += normalNumberMap[string(s[0])]
				s = s[1:]
			}
		} else {
			result += normalNumberMap[s]
			s = s[1:]
		}
	}

	return result
}

//func main() {
//	fmt.Println(romanToInt("III"))
//	fmt.Println(romanToInt("LVIII"))
//}
