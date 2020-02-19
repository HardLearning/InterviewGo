package main

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
    for i := 0; i < len(nums); i++ {
    	r, ok := mp[nums[i]]
    	if ok {
    		val := []int{r, i}
    		return val
		} else {
			mp[target - nums[i]] = i
		}
	}
	res := []int{0, 0}
	return res
}


//func main() {
//	nums := []int{0, 1 , 2, 3, 4, 5, 6}
//	fmt.Println(twoSum(nums, 11))
//}