package day

func DistinctDifferenceArray(nums []int) []int {
	res := make([]int, 0)
	preMap := make(map[int]struct{})

	for i, num := range nums {
		preMap[num] = struct{}{}
		sufMap := make(map[int]struct{})
		for _, t := range nums[i+1:] {
			sufMap[t] = struct{}{}
		}
		res = append(res, len(preMap)-len(sufMap))
	}
	return res
}
