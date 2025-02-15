package main

import (
	"sort"
)

// original version，写的多少有点丑陋了
func threeSum1(nums []int) [][]int {
	res := make([][]int, 0)
	sortMap := make(map[[3]int]struct{})
	for i := 0; i < len(nums)-2; i++ {
		tmp := make(map[int]struct{})
		for j := i + 1; j < len(nums); j++ {
			if _, ok := tmp[-nums[i]-nums[j]]; ok {
				val := []int{nums[i], nums[j], -nums[i] - nums[j]}
				sort.Ints(val)
				if _, ok := sortMap[[3]int{val[0], val[1], val[2]}]; ok {
					continue
				} else {
					sortMap[[3]int{val[0], val[1], val[2]}] = struct{}{}
				}
				res = append(res, val)
			} else {
				tmp[nums[j]] = struct{}{}
			}
		}
	}
	return res
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[l] + nums[r]
			if sum == target {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if sum < target {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
