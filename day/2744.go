package day

func MaximumNumberOfStringPairs(words []string) int {
	originMap := make(map[string]struct{})
	count := 0
	for _, word := range words {
		originMap[word] = struct{}{}
	}
	for _, word := range words {
		delete(originMap, word)
		bytes := []byte(word)
		// 使用双指针进行反转
		for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
			bytes[i], bytes[j] = bytes[j], bytes[i]
		}
		str := string(bytes)
		if _, ok := originMap[str]; ok {
			count++
			delete(originMap, str)
		}
	}
	return count
}
