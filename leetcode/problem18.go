package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	res := make([][]int, 0)
	if length < 4 {
		return res
	}
	sort.Ints(nums)

	var sum int
	for a:=0; a < length - 3; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		for b:=a+1; b < length - 2; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}

			for c,d := b+1, length-1; c<d; {
				sum = nums[a] + nums[b] + nums[c] + nums[d]
				switch {
				case sum > target:
					d--
					for c<d && nums[d] == nums[d+1] {d--}
				case sum < target:
					c++
					for c<d && nums[c] == nums[c-1] {c++}
				case sum == target:
					res = append(res, []int{nums[a], nums[b], nums[c], nums[d]} )
					d--
					for c<d && nums[d] == nums[d+1] {d--}
					c++
					for c<d && nums[c] == nums[c-1] {c++}
				}
			}
		}
	}

	return res
}