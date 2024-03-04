package hot

import (
	"leetcode/common"
	"math"
)

func isValidBST(root *common.TreeNode) bool {
	return helper98(root, math.MinInt64, math.MaxInt64)
}

func helper98(root *common.TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return helper98(root.Left, min, root.Val) && helper98(root.Right, root.Val, max)
}
