package _065

func minOperations(nums []int, k int) int {
	res := 0
	for _, num := range nums {
		if num < k {
			res += 1
		}
	}
	return res
}
