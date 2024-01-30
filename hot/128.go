package hot

// 100, 4, 200, 1, 3, 2
func LongestConsecutive(nums []int) int {
	res := 0
	numMap := make(map[int]bool)
	for _, num := range nums {
		numMap[num] = true
	}
	for key := range numMap {
		if !numMap[key-1] {
			tmpKey := key
			count := 1
			for numMap[tmpKey+1] {
				tmpKey += 1
				count++
			}
			if count > res {
				res = count
			}
		}
	}
	return res
}
