package hot

import (
	"leetcode/common"
	"testing"
)

func Test437(t *testing.T) {
	//println(pathSum(common.BuildTreeFromSlice([]interface{}{10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1}), 8))
	println(pathSum(common.BuildTreeFromSlice([]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1}), 22))
}
