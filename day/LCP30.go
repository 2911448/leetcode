package day

import "sort"

// 100, 100, 100, -250, -60, -140, -50, -50, 100, 150
func MagicTower(nums []int) int {
	count, hp := 0, 1
	hurt := make([]int, 0)
	moveItem := 0
	for _, num := range nums {
		hp += num
		if num < 0 {
			hurt = append(hurt, num)
			if hp <= 0 {
				sort.Ints(hurt)

				max := hurt[0]
				hp -= max
				moveItem += max

				hurt = hurt[1:]
				count++
			}

		}
	}
	if hp+moveItem < 0 {
		return -1
	}
	return count
}
