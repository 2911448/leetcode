package hot

func Merge(intervals [][]int) [][]int {
	result := make([][]int, 0)
	length := len(intervals)
	if length == 0 {
		return result
	}
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if intervals[i][0] > intervals[j][0] {
				tmp := intervals[i]
				intervals[i] = intervals[j]
				intervals[j] = tmp
			}
		}
	}

	result = append(result, intervals[0])
	for i := 1; i < length; i++ {
		resultLen := len(result)
		if intervals[i][0] > result[resultLen-1][1] || intervals[i][1] < result[resultLen-1][0] {
			result = append(result, intervals[i])
		} else {
			tmp := result[resultLen-1]
			if intervals[i][1] > tmp[1] {
				tmp[1] = intervals[i][1]
			}
			if intervals[i][0] < tmp[0] {
				tmp[0] = intervals[i][0]
			}
		}
	}
	return result
}
