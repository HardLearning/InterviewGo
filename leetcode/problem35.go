package main

/**
* @author moorewu
* @E-mail moorewu@tencent.com
* @date 2020/2/27 9:41 上午
 */

func searchInsert(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	result := length
	left, right := 0, length-1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
			result = mid
		}
	}
	return result
}

//func main() {
//	nums1 := []int{1, 3, 5, 6}
//	fmt.Println(searchInsert(nums1, 5))
//
//	nums2 := []int{1, 3, 5, 6}
//	fmt.Println(searchInsert(nums2, 2))
//
//	nums3 := []int{1, 3, 5, 6}
//	fmt.Println(searchInsert(nums3, 7))
//
//	nums4 := []int{1, 3, 5, 6}
//	fmt.Println(searchInsert(nums4, 0))
//}
