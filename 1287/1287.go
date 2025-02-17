package main

func findSpecialInteger(arr []int) int {
	times := len(arr)/4 + 1
	curTimes := 1
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			curTimes++
			if curTimes >= times {
				res = arr[i]
				break
			}
		} else {
			curTimes = 1
		}
	}
	return res
}
