package hot

// 1, 2, 3		3
// 1, 1, 1 		2
func SubarraySum(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		preSum := 0
		for right := i; right >= 0; right-- {
			preSum += nums[right]
			if preSum == k {
				res++
			}
		}
	}
	return res
}
