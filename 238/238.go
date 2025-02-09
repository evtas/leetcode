package main

func productExceptSelf(nums []int) []int {
	l, r := make([]int, len(nums)), make([]int, len(nums))
	l[0], l[1] = 1, nums[0]
	r[len(nums)-1], r[len(nums)-2] = 1, nums[len(nums)-1]
	for i := 2; i < len(nums); i++ {
		l[i] = nums[i-1] * l[i-1]
	}
	for i := len(nums) - 3; i >= 0; i-- {
		r[i] = nums[i+1] * r[i+1]
	}
	res := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		res = append(res, l[i]*r[i])
	}
	return res
}

func main() {
	res := productExceptSelf([]int{1, 2, 3, 4})
	for _, num := range res {
		println(num)
	}
}
