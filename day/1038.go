package day

import "leetcode/common"

func BstToGst(root *common.TreeNode) *common.TreeNode {
	sum := 0
	var dfs func(*common.TreeNode)
	dfs = func(node *common.TreeNode) {
		if node != nil {
			dfs(node.Right)
			sum += node.Val
			node.Val = sum
			dfs(node.Left)
		}
	}
	dfs(root)
	return root
}
