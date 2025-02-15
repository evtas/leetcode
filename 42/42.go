package main

import "fmt"

// 写的有点丑陋，但是通过了，双指针，左右各一遍解决
func trap(height []int) int {
	res := 0
	if len(height) < 1 {
		return 0
	}
	leftPtr := 0
	for leftPtr < len(height)-1 && height[leftPtr] < height[leftPtr+1] {
		leftPtr++
	}
	rightPtr := leftPtr + 1
	if rightPtr > len(height)-1 {
		return 0
	}
	for leftPtr >= 0 && rightPtr < len(height) && leftPtr < rightPtr {
		for rightPtr < len(height) && height[rightPtr] < height[leftPtr] {
			rightPtr++
		}
		if leftPtr < 0 || rightPtr > len(height)-1 {
			continue
		}
		leftVal := min(height[leftPtr], height[rightPtr])
		for i := leftPtr + 1; i < rightPtr; i++ {
			res += leftVal - height[i]
			fmt.Println(leftPtr, rightPtr, i, leftVal, height[i], res)
		}
		leftPtr = rightPtr
		rightPtr++
	}
	flag := leftPtr
	fmt.Println(res)
	if leftPtr >= len(height)-1 {
		return res
	}

	rightPtr = len(height) - 1
	for rightPtr > 0 && height[rightPtr] < height[rightPtr-1] {
		rightPtr--
	}
	leftPtr = rightPtr - 1
	for leftPtr >= 0 && rightPtr < len(height) && leftPtr >= flag {
		for leftPtr >= 0 && height[leftPtr] < height[rightPtr] {
			leftPtr--
		}
		if leftPtr < 0 || rightPtr > len(height)-1 {
			continue
		}
		rightVal := min(height[rightPtr], height[leftPtr])
		for i := rightPtr - 1; i > leftPtr; i-- {
			res += rightVal - height[i]
		}
		rightPtr = leftPtr
		leftPtr--
	}
	return res
}
