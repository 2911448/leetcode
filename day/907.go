package day

func SumSubarrayMins(arr []int) int {
	const mod int = 1e9 + 7
	sum := 0
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			sum += getMin(arr[i:j+1]) % mod
		}
	}
	return sum
}

func getMin(arr []int) int {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}
