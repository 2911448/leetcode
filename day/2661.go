package day

// []int{1, 3, 4, 2}, [][]int{{1, 4}, {2, 3}}
func FirstCompleteIndex(arr []int, mat [][]int) int {
	n, m := len(mat), len(mat[0])
	// 转换成 MAP
	matMap := make(map[int][2]int)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matMap[mat[i][j]] = [2]int{i, j}
		}
	}
	rowCnt, colCnt := make([]int, n), make([]int, m)
	for i := 0; i < n; i++ {
		rowCnt[i] = 0
	}
	for j := 0; j < m; j++ {
		colCnt[j] = 0
	}
	for i := 0; i < len(arr); i++ {
		v := matMap[arr[i]]
		rowCnt[v[0]]++
		if rowCnt[v[0]] == m {
			return i
		}
		colCnt[v[1]]++
		if colCnt[v[1]] == n {
			return i
		}
	}

	return -1
}
