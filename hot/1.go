package hot

func TwoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, num := range nums {
		mp[num] = i
	}
	for i, num := range nums {
		if v, ok := mp[target-num]; ok {
			if v != i {
				return []int{i, v}
			}
		}
	}
	return []int{}
}
