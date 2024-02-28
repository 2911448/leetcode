package hot

import "leetcode/common"

// 1->2->3->4->5->nil
func ReverseKGroup(head *common.ListNode, k int) *common.ListNode {
	var res *common.ListNode
	current := head
	count := 0
	left := head
	for current != nil {
		count++
		if count == k {
			next := current.Next
			current.Next = nil
			res = joinList(res, ReverseList(left))
			current = next
			left = next
			count = 0
		} else {
			current = current.Next
		}
	}
	if count > 0 {
		res = joinList(res, left)
	}
	return res
}
func joinList(list1, list2 *common.ListNode) *common.ListNode {
	if list1 == nil {
		return list2
	}
	res, l2 := &common.ListNode{Next: list1}, list2
	current, tail := res, res
	for current.Next != nil {
		current = current.Next
		tail = current
	}
	for l2 != nil {
		tail.Next = l2
		tail = tail.Next
		l2 = l2.Next
	}
	return res.Next
}

func ReverseKGroup1(head *common.ListNode, k int) *common.ListNode {
	if head == nil || head.Next == nil || k <= 0 {
		return head
	}

	dummy := &common.ListNode{Val: -1, Next: head}
	pre := dummy

	for cur := head; cur != nil; cur = cur.Next {
		tmp := cur
		for i := 0; i < k-1 && tmp != nil; i++ {
			tmp = tmp.Next
		}
		if tmp == nil {
			break
		}
		for i := 0; i < k-1; i++ {
			nextNode := cur.Next
			cur.Next = nextNode.Next
			nextNode.Next = pre.Next
			pre.Next = nextNode
		}
		pre = cur
	}
	return dummy.Next
}
