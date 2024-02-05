package main

import "leetcode/hot"

func main() {
	// 两数之和 1
	//for _, i := range hot.TwoSum([]int{3, 2, 4}, 6) {
	//	println(i)
	//}

	// 两数相加
	//println(hot.AddTwoNumbers(common.MakeListNode([]int{9, 9, 9, 9, 9, 9, 9}), common.MakeListNode([]int{9, 9, 9, 9})).PrintListNode())

	// 无重复字符串的最长子串
	//println(hot.LengthOfLongestSubstring("abcabcbb"))

	// 最长回文子串 *******
	//println(hot.LongestPalindrome("babad"))

	// 盛水最多的容器
	//println(hot.MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

	// 删除链表的倒数第 N 个节点 删除
	//println(hot.RemoveNthFromEnd(common.MakeListNode([]int{1, 2, 3, 4, 5}), 2).PrintListNode())

	// 最长连续序列
	//println(hot.LongestConsecutive([]int{100, 4, 200, 1, 3, 2}))

	// 接雨水
	//println(hot.Trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	//println(hot.Trap([]int{4, 2, 0, 3, 2, 5}))

	// 找到字符串中所有字母异位词
	//for _, str := range hot.FindAnagrams("cbaebabacd", "abc") {
	//	println(str)
	//}

	// 和为 K 的子数组
	//print(hot.SubarraySum([]int{1, 2, 3}, 3))
	//print(hot.SubarraySum([]int{1, 1, 1}, 2))

	// 合并区间
	//for _, ints := range hot.Merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}) {
	//	for _, i := range ints {
	//		print(i)
	//	}
	//	println("--")
	//}
	for _, ints := range hot.Merge([][]int{{1, 4}, {0, 0}}) {
		for _, i := range ints {
			print(i)
		}
		println("--")
	}
}
