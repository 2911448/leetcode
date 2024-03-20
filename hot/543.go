package hot

import (
	"leetcode/common"
	"leetcode/tool"
)

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
	return tool.Max(left, right) + 1
}
