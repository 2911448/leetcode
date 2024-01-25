package day

import "strconv"

// 计算 K 位置下标对应元素的和
// 5, 10, 1, 5, 2 ; k = 1
func SumIndicesWithKSetBits(nums []int, k int) int {
	res := 0
	count := func(str string) int {
		sum := 0
		for i := range str {
			if str[i] == '1' {
				sum++
			}
		}
		return sum
	}
	for i, num := range nums {
		binaryStr := strconv.FormatInt(int64(i), 2)
		if count(binaryStr) == k {
			res += num
		}
	}
	return res
}
