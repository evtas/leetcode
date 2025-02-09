package main

import "fmt"

func removeElement(nums []int, val int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			cnt++
		}
	}
	if cnt == 0 {
		return 0
	}
	if cnt == len(nums) {
		return cnt
	}
	p1, p2 := 0, len(nums)-1
	for p1 < p2 {
		for nums[p1] != val {
			p1++
		}
		for nums[p2] == val {
			p2--
		}
		if p1 < p2 {
			nums[p1], nums[p2] = nums[p2], nums[p1]
		}
	}
	return cnt
}

func main() {
	nums := []int{3, 3}
	val := 5
	fmt.Println(removeElement(nums, val))
}
