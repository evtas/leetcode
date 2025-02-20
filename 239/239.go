package main

// 单调队列
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 1, len(nums)-k+1)
	q := make([]int, 0)
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	for i := 0; i < k; i++ {
		push(i)
	}

	res[0] = nums[q[0]]
	for i := k; i < len(nums); i++ {
		push(i)
		for q[0] < i-k+1 {
			q = q[1:]
		}
		res = append(res, nums[q[0]])
	}
	return res
}
