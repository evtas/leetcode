package main

func jump(nums []int) int {
	maxFar := 0
	end := 0
	step := 0
	for i := 0; i < len(nums)-1; i++ {
		maxFar = max(maxFar, i+nums[i])
		if i == end {
			end = maxFar
			step += 1
		}
	}
	return step
}
