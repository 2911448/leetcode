package hot

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	index := -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			index = mid
			break
		}
	}
	if index != -1 {
		tmpLeft := index
		for tmpLeft > 0 {
			tl := tmpLeft - 1
			if nums[tl] == target {
				tmpLeft--
			} else {
				break
			}
		}
		tmpRight := index
		for tmpRight < len(nums)-1 {
			tr := tmpRight + 1
			if nums[tr] == target {
				tmpRight++
			} else {
				break
			}

		}
		return []int{tmpLeft, tmpRight}
	} else {
		return []int{-1, -1}
	}
}
