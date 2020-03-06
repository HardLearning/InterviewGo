package main

import "sort"

/**
* @author moorewu
* @E-mail moorewu@tencent.com
* @date 2020/3/6 10:47 上午
 */

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
