package main

import "fmt"

// 超时
// func maxArea(height []int) int {
// 	tmp := make([]int, len(height))
// 	for i := 1; i < len(height); i++ {
// 		for j := 0; j < i; j++ {
// 			tmp[j] = max(min(height[i], height[j])*(i-j), tmp[j])
// 		}
// 	}
// 	res := 0
// 	for _, val := range tmp {
// 		if val > res {
// 			res = val
// 		}
// 	}
// 	return res
// }

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res := min(height[left], height[right]) * (right - left)
	for left < right {
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
		res = max(res, min(height[left], height[right])*(right-left))
	}
	return res
}

func main() {
	res := maxArea([]int{8, 7, 2, 1})
	fmt.Printf("res: %v\n", res)
}
