package main

import (
	"leetcode/hot"
)

func main() {
	// 两数之和 1
	//for _, i := range hot.TwoSum([]int{3, 2, 4}, 6) {
	//	println(i)
	//}

	// 两数相加
	//println(hot.AddTwoNumbers(common.MakeListNode([]int{9, 9, 9, 9, 9, 9, 9}), common.MakeListNode([]int{9, 9, 9, 9})).PrintListNode())

	// 无重复字符串的最长子串
	println(hot.LengthOfLongestSubstring("abcabcbb"))

}
