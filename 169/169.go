package _69

func majorityElement(nums []int) int {
	candidate := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		switch {
		case count == 0:
			candidate = nums[i]
			count++
		case nums[i] == candidate:
			count++
		case nums[i] != candidate:
			count--
		default:
		}
	}
	return candidate
}

func main() {

}
