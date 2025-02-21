package main

import "fmt"

func maxSubArray(nums []int) int {
	preSum := make([]int, len(nums))
	preSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i]
	}
	left, right := 0, 1
	fmt.Println(preSum)
	maxSum := preSum[left]
	for right < len(nums) {
		maxSum = max(maxSum, preSum[right]-preSum[left])
		maxSum = max(maxSum, preSum[right])
		if preSum[right] < preSum[left] {
			left = right
		}
		right++
	}
	return max(maxSum, preSum[len(preSum)-1])
}
