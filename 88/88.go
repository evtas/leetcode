package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for nums1Ptr, nums2Ptr, ptr := m-1, n-1, m+n-1; nums1Ptr >= 0 || nums2Ptr >= 0; ptr-- {
		var cur int
		if nums1Ptr == -1 {
			cur = nums2[nums2Ptr]
			nums2Ptr--
		} else if nums2Ptr == -1 {
			cur = nums1[nums1Ptr]
			nums1Ptr--
		} else if nums1[nums1Ptr] > nums2[nums2Ptr] {
			cur = nums1[nums1Ptr]
			nums1Ptr--
		} else {
			cur = nums2[nums2Ptr]
			nums2Ptr--
		}
		nums1[ptr] = cur
	}
	fmt.Printf("%+v\n", nums1)
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
}
