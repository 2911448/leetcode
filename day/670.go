package day

import "strconv"

func MaximumSwap(num int) int {
	f := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	max := num
	str := []byte(strconv.Itoa(num))
	for i := range str {
		for j := range str[:i] {
			str[i], str[j] = str[j], str[i]
			tmp, _ := strconv.Atoi(string(str))
			max = f(max, tmp)
			str[j], str[i] = str[i], str[j]
		}
	}
	return max
}
