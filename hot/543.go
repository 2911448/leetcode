package hot

import "leetcode/common"

func diameterOfBinaryTree(root *common.TreeNode) int {
	left := helper(root.Left)
	right := helper(root.Right)
	return left + right
}
func helper(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left)
	right := helper(root.Right)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
