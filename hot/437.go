package hot

import "leetcode/common"

func pathSum(root *common.TreeNode, targetSum int) int {
	var res int
	var dfs func(root *common.TreeNode, targetSum int)
	dfs = func(root *common.TreeNode, targetSum int) {
		if root == nil {
			return
		}
		targetSum -= root.Val
		dfs(root.Left, targetSum)
		dfs(root.Right, targetSum)
		if targetSum == 0 {
			res++
		}
		targetSum += root.Val
	}
	target := targetSum
	dfs(root.Left, target)
	dfs(root.Right, target)
	return res
}
