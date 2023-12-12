package day

import (
	"reflect"
	"sort"
)

func CloseStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	var w1Mask, w2Mask int
	var sCnt, tCnt [26]int
	for _, c := range word1 {
		w1Mask |= 1 << (c - 'a')
		sCnt[c-'a']++
	}
	for _, c := range word2 {
		w2Mask |= 1 << (c - 'a')
		tCnt[c-'a']++
	}
	sort.Ints(sCnt[:])
	sort.Ints(tCnt[:])
	return w1Mask == w2Mask && reflect.DeepEqual(sCnt[:], tCnt[:])

}
