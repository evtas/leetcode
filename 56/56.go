package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	curVal := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= curVal[1] {
			curVal[1] = max(curVal[1], intervals[i][1])
		} else {
			res = append(res, curVal)
			curVal = intervals[i]
		}
	}
	res = append(res, curVal)
	return res
}

func main() {
	intervals := [][]int{
		{1, 3},
		{8, 10},
		{2, 6},
		{15, 18},
	}
	res := merge(intervals)
	fmt.Println(res)
}
