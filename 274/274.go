package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	sort.Ints(citations)
	h := 0
	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] > h && len(citations)-i > h {
			h++
		}
	}
	return h
}

func main() {
	rr := hIndex([]int{1, 3, 1})
	fmt.Println(rr)
}
