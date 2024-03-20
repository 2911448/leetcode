package hot

import "testing"

func TestTrap(t *testing.T) {
	println(Trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
