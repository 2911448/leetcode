package hot

func MaxArea(height []int) int {
	res := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	length := len(height)
	left, right := 0, length-1
	for left < right {
		if height[left] < height[right] {
			res = max(res, (right-left)*height[left])
			left++
		} else {
			res = max(res, (right-left)*height[right])
			right--
		}
	}
	return res
}
