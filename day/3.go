package day

import "leetcode/tool"

func LengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	right, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for right+1 < n && m[s[right+1]] == 0 {
			// 不断地移动右指针
			m[s[right+1]]++
			right++
		}
		ans = tool.Max(ans, right-i+1)
	}
	return ans
}
