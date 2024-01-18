package day

import (
	"math"
	"sort"
)

//func MinimumRemoval(beans []int) int64 {
//	sort.Slice(beans, func(i, j int) bool {
//		return beans[i] < beans[j]
//	})
//	var count int64
//	if len(beans) <= 1 {
//		return count
//	}
//	for right := len(beans) - 1; right > 0; right-- {
//		sum := 0
//		for _, i := range beans[:right] {
//			sum += i
//		}
//		if sum >= beans[right] {
//			continue
//		} else {
//			preSum := 0
//			for _, num := range beans[:right] {
//				preSum += num
//			}
//			sufSum := 0
//			for _, num := range beans[right:] {
//				sufSum += num
//			}
//			count = int64(preSum + (sufSum - beans[right]*len(beans[right:])))
//		}
//	}
//
//	return count
//}

func MinimumRemoval(beans []int) int64 {
	min := func(a, b int64) int64 {
		if a > b {
			return b
		}
		return a
	}
	sort.Slice(beans, func(i, j int) bool {
		return beans[i] < beans[j]
	})
	var count int64 = math.MaxInt64
	if len(beans) <= 1 {
		return 0
	}
	preSum := 0
	sufSum := 0
	allEqual := true
	for _, num := range beans {
		sufSum += num
		if num != beans[0] {
			allEqual = false
		}
	}
	if allEqual {
		return 0
	}
	for right := 0; right < len(beans); right++ {
		count = min(count, int64(preSum+(sufSum-beans[right]*len(beans[right:]))))
		preSum += beans[right]
		sufSum -= beans[right]
	}

	return count
}
