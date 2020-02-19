package main

func removeELement(nums []int, val int) int {
	len := len(nums)
	if len == 0 {
		return 0
	}
	fast := 0
	slow := 0
	for fast < len {
		if nums[fast] == val {
			fast++
		} else {
			nums[slow] = nums[fast]
			fast++
			slow++
		}
	}
	return len - (fast - slow)
}


//func main() {
//	nums := []int{3,2,2,3}
//	rest := removeELement(nums, 3)
//	fmt.Println(rest)
//	fmt.Print(nums)
//}