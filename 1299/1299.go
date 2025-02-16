package main

func replaceElements(arr []int) []int {
	curMax := arr[len(arr)-1]
	res := make([]int, len(arr))
	res[len(arr)-1] = -1
	for i := len(arr) - 2; i >= 0; i-- {
		curMax = max(curMax, arr[i+1])
		res[i] = curMax
	}
	return res
}
