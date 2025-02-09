package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	checkPoints := make([]int, 0)
	for i := 0; i < len(gas); i++ {
		if gas[i] >= cost[i] {
			checkPoints = append(checkPoints, i)
		}
	}
	for i := 0; i < len(checkPoints); i++ {
		partial := checkPoints[i]
		restFeul := 0
		canPass := true
		for j := 0; j < len(gas); j++ {
			restFeul = restFeul + gas[(j+partial)%len(gas)] - cost[(j+partial)%len(gas)]
			if restFeul < 0 {
				canPass = false
				break
			}
		}
		if canPass {
			return checkPoints[i]
		}
	}
	return -1
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	result := canCompleteCircuit(gas, cost)
	fmt.Println(result)
}
