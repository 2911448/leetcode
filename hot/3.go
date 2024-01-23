package hot

// abcabcbb
// pwwkew
// bbbbb
// tmmzuxt
func LengthOfLongestSubstring(s string) int {
	bytes := []byte(s)
	mp := make(map[byte]int)
	max, left := 0, 0
	f := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	length := len(bytes)
	for i := 0; i < length; i++ {
		if v, ok := mp[bytes[i]]; ok {
			if v >= left {
				left = f(left, v) + 1
			}
		}
		mp[bytes[i]] = i
		max = f(max, i-left+1)
	}
	return max
}
