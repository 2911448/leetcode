package day

func AlternatingSubarray(nums []int) int {
	f := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	res := -1
	length := len(nums)
	left := 0
	for i := 1; i < length; i++ {
		l := i - left + 1
		if nums[i]-nums[left] == (l-1)%2 {
			res = f(res, l)
		} else {
			if nums[i]-nums[i-1] == 1 {
				left = i - 1
				res = f(res, 2)
			} else {
				left = i
			}

		}
	}
	return res
}
