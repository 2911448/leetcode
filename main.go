package main

import (
	"fmt"
	"leetcode/common"
	"leetcode/day"
)

func main() {

	//1410
	//fmt.Println(day.EntityParser("&amp; is an HTML entity but &ambassador; is not."))

	//2
	//fmt.Println(day.AddTwoNumbers(common.MakeListNode([]int{0, 8, 6, 5, 6, 8, 3, 5, 7}), common.MakeListNode([]int{6, 7, 8, 0, 8, 5, 8, 9, 7})).PrintListNode())

	//2824
	//fmt.Println(day.CountPairs([]int{-6, 2, 5, -2, -7, -1, 3}, -2))
	//fmt.Println(day.CountPairs2([]int{-6, 2, 5, -2, -7, -1, 3}, -2))

	//3
	//fmt.Println(day.LengthOfLongestSubstring("bbtablud"))

	//907
	//fmt.Println(day.SumSubarrayMins([]int{11, 81, 94, 43, 3}))

	//1657
	//fmt.Println(day.CloseStrings("abc", "bca"))

	//2661
	//fmt.Println(day.FirstCompleteIndex([]int{1, 3, 4, 2}, [][]int{{1, 4}, {2, 3}}))

	//1038
	//tree := common.BuildTreeFromSlice([]interface{}{4, 1, 6, 0, 2, 5, 7, nil, nil, nil, 3, nil, nil, nil, 8})
	//common.PrintTree(day.BstToGst(tree))

	//162
	//	fmt.Println(day.FindPeakElement([]int{1, 2}))

	// 2696
	//fmt.Println(day.MinLength("ABFCACDB"))

	// 82
	fmt.Println(day.DeleteDuplicates(common.MakeListNode([]int{1, 2, 3, 3, 4, 4, 5})).PrintListNode())
}
