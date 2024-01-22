package hot

import "leetcode/common"

func AddTwoNumbers(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	res := &common.ListNode{}
	tmp := res
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			carry += l1.Val
		}
		if l2 != nil {
			carry += l2.Val
		}

		tmp.Next = &common.ListNode{Val: carry % 10}
		carry /= 10
		tmp = tmp.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return res.Next
}
