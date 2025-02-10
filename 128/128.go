package main

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	longest := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentLongest := 1
			currentNum := num
			for numSet[currentNum+1] {
				currentLongest++
				currentNum++
			}
			if longest < currentLongest {
				longest = currentLongest
			}
		}
	}
	return longest
}

func main() {
	res := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	println(res)
}
