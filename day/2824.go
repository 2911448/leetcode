package day

import "sort"

func CountPairs(nums []int, target int) int {
	res := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] < target {
				res++
			}
		}
	}
	return res
}

func CountPairs2(nums []int, target int) int {
	res := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		j := i + 1
		for j < len(nums) {
			if nums[i]+nums[j] < target {
				res++
			} else {
				break
			}
			j++
		}
	}
	return res
}
