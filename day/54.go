package day

func SpiralOrder(matrix [][]int) []int {
	const visit = -200
	direction := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	row, col := len(matrix), len(matrix[0])
	directionIndex, left, right, total := 0, 0, 0, row*col
	res := make([]int, 0, total)
	for len(res) < total {
		res = append(res, matrix[left][right])
		matrix[left][right] = visit
		nextLeft, nextRight := left+direction[directionIndex][0], right+direction[directionIndex][1]
		if nextLeft < 0 || nextLeft == row || nextRight < 0 || nextRight == col || matrix[nextLeft][nextRight] == visit {
			directionIndex = (directionIndex + 1) % 4
		}
		left += direction[directionIndex][0]
		right += direction[directionIndex][1]
	}
	return res
}
