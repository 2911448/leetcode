package hot

import (
	"leetcode/common"
	"testing"
)

func Test543(t *testing.T) {
	println(diameterOfBinaryTree(common.BuildTreeFromSlice([]interface{}{1, 2, 3, 4, 5})))
	println(diameterOfBinaryTree(common.BuildTreeFromSlice([]interface{}{1, 2, nil})))
}
