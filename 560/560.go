package main

// 区间和=前缀和的差值
func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	mp := make(map[int]int)
	mp[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := mp[pre-k]; ok {
			count += mp[pre-k]
		}
		mp[pre] += 1
	}
	return count
}
