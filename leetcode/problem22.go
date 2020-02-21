package main

func generateParenthesis(n int) []string {
	var ret []string
	str := ""
	dfs(n, n, str, &ret)
	return ret
}

//递归解题
//n1代表待插入的左括号个数 n2代表待插入的右括号个数
func dfs(n1, n2 int, str string, ret *[]string) {
	if n1 == 0 && n2 == 0 {
		*ret = append(*ret, str)
		return
	}
	if n1 <= n2 && n1 > 0 {
		dfs(n1-1, n2, str+"(", ret)
	}
	if n2 > n1 {
		dfs(n1, n2-1, str+")", ret)
	}
}

//func main() {
//	ret := generateParenthesis(3)
//	fmt.Println(ret)
//}
