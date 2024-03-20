package hot

import "leetcode/tool"

func Trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	result := 0
	for left < right {
		leftMax = tool.Max(leftMax, height[left])
		rightMax = tool.Max(rightMax, height[right])
		if height[left] < height[right] {
			result += leftMax - height[left]
			left++
		} else {
			result += rightMax - height[right]
			right--
		}
	}
	return result
}
