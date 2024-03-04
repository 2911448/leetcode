package hot

import "leetcode/common"

func rightSideView(root *common.TreeNode) []int {
	result := make([][]int, 0)
	queue := []*common.TreeNode{root}
	count := 0
	for len(queue) > 0 {
		result = append(result, []int{})
		tmpQueue := []*common.TreeNode{}
		for i := 0; i < len(queue); i++ {
			node := queue[i]
			result[count] = append(result[count], node.Val)
			if node.Left != nil {
				tmpQueue = append(tmpQueue, node.Left)
			}
			if node.Right != nil {
				tmpQueue = append(tmpQueue, node.Right)
			}
		}
		count++
		queue = tmpQueue
	}
	res := make([]int, len(result))
	for i, ints := range result {
		res[i] = ints[len(ints)-1]
	}
	return res
}
