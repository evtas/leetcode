package main

import "fmt"

func canJump(nums []int) bool {
	maxPos := 0
	for i := 0; i < len(nums)-1; i++ {
		if i <= maxPos {
			maxPos = max(maxPos, i+nums[i])
		}
	}
	return maxPos >= len(nums)-1
}

func main() {
	fmt.Println(canJump([]int{2, 5, 0, 0}))
}
