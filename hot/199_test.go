package hot

import (
	"leetcode/common"
	"testing"
)

func Test199(t *testing.T) {
	for _, i := range rightSideView(common.BuildTreeFromSlice([]interface{}{1, 2, nil})) {
		println(i)
	}
}
