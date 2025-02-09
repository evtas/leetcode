package main

import "fmt"

// 滑动窗口的思想

func removeDuplicates(nums []int) int {
	k := 2
	if len(nums) < k {
		return len(nums)
	}
	slow, fast := k, k
	for fast < len(nums) {
		if nums[fast] != nums[slow-k] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
