package main

import (
	"math"
	"sort"
)

/**
* @author moorewu
* @E-mail moorewu@tencent.com
* @date 2020/2/27 9:53 上午
 */

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	min := nums[0] + nums[1] + nums[2]
	max := nums[n-3] + nums[n-2] + nums[n-1]
	if target > max {
		return max
	}
	if target < min {
		return min
	}

	result := min
	for i := 0; i < n; i++ {
		j, k := i+1, n-1
		for j < k {
			temp := nums[i] + nums[j] + nums[k]
			if math.Abs(float64(target-result)) > math.Abs(float64(target-temp)) {
				result = temp
			}
			if temp < target {
				j++
			} else if temp > target {
				k--
			} else {
				return result
			}
		}
	}
	return result
}

//func main() {
//	nums := []int{-1, 2, 1, -4}
//	fmt.Println(threeSumClosest(nums, 1))
//}
