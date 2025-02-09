package main

import "fmt"

func removeDuplicates(nums []int) int {
	left, right := 0, 1
	for right < len(nums) {
		if nums[right] != nums[left] {
			left++
			nums[left] = nums[right]
		}
		right++
	}
	return left + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}
