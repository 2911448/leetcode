package hot

import "leetcode/common"

func RemoveNthFromEnd(head *common.ListNode, n int) *common.ListNode {
	nodeList := make([]*common.ListNode, 0)
	dummy := &common.ListNode{Next: head}
	for node := dummy; node != nil; node = node.Next {
		nodeList = append(nodeList, node)
	}
	pre := nodeList[len(nodeList)-n-1]
	pre.Next = pre.Next.Next
	return dummy.Next
}
