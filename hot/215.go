package hot

func findKthLargest(nums []int, k int) int {
	buildHeap(nums)
	for i := 0; i < k-1; i++ {
		pop(&nums)
	}
	return nums[0]
}

func buildHeap(nums []int) {
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		down(nums, i, n)
	}
}

func down(nums []int, i, n int) {
	left := 2*i + 1
	if left >= n || left < 1 {
		return
	}
	max := left
	right := 2*i + 2
	if right < n && nums[right] > nums[left] {
		max = right
	}
	// 如果左右子节点都小于当前节点退出循环
	if nums[max] < nums[i] {
		return
	}
	// 交换节点
	nums[i], nums[max] = nums[max], nums[i]
	// 递归 max为根的子树满不满足堆的条件
	down(nums, max, n)
}

func pop(nums *[]int) int {
	last := len(*nums) - 1
	// 将根节点和尾节点交换
	(*nums)[0], (*nums)[last] = (*nums)[last], (*nums)[0]
	// 重构堆
	down(*nums, 0, last)
	// 舍去最小的
	rst := (*nums)[last]
	(*nums) = (*nums)[:last]
	return rst
}
