package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	var temp = x
	var res = 0
	for ; temp > 0; temp = temp / 10 {
		res = res*10 + temp%10
	}
	return res == x
}

//func main() {
//	fmt.Println(isPalindrome(-100))
//	fmt.Println(isPalindrome(100))
//	fmt.Println(isPalindrome(101))
//}
