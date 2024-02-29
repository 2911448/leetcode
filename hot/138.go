package hot

import "leetcode/common"

func CopyRandomList(head *common.RandomNode) *common.RandomNode {
	cache := make(map[*common.RandomNode]*common.RandomNode)
	var get func(node *common.RandomNode) *common.RandomNode
	get = func(node *common.RandomNode) *common.RandomNode {
		if node == nil {
			return nil
		}
		if v, ok := cache[node]; ok {
			return v
		}
		newNode := &common.RandomNode{Val: node.Val}
		cache[node] = newNode
		newNode.Next = get(node.Next)
		newNode.Random = get(node.Random)
		return newNode
	}
	return get(head)
}
