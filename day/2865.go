package day

// 1,5,2,5,6,4,6,3,4,5
func MaximumSumOfHeights(maxHeights []int) int64 {
	length := len(maxHeights)
	var res int64
	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	for i := 0; i < length; i++ {
		curVal := maxHeights[i]
		sum := curVal
		for j := i - 1; j >= 0; j-- {
			curVal = min(curVal, maxHeights[j])
			sum += curVal
		}

		curVal = maxHeights[i]
		for j := i + 1; j < length; j++ {
			curVal = min(curVal, maxHeights[j])
			sum += curVal
		}
		res = max(res, int64(sum))
	}

	return res
}
