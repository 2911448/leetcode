package day

import "leetcode/common"

func DeleteDuplicates(head *common.ListNode) *common.ListNode {
	if head == nil {
		return nil
	}
	res := &common.ListNode{Next: head}
	cur := res
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			v := cur.Next.Val
			for cur.Next.Next != nil && cur.Next.Next.Val == v {
				cur.Next = cur.Next.Next
			}
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return res.Next
}
