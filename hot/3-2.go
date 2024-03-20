package hot

import (
	"leetcode/tool"
)

func lengthOfLongestSubstring(s string) int {
	mp := make(map[byte]int)
	n := len(s)
	right, res := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(mp, s[i-1])
		}
		for right+1 < n && mp[s[right+1]] == 0 {
			mp[s[right+1]]++
			right++
		}
		res = tool.Max(res, right-i+1)
	}
	return res
}
