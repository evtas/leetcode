package main

func lengthOfLongestSubstring1(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	res := 0
	tmp := make(map[byte]int)
	left, right := 0, 1
	tmp[s[left]] = 0
	for right < len(s) {
		if _, ok := tmp[s[right]]; ok {
			res = max(res, len(tmp))
			left = tmp[s[right]] + 1
			right = left + 1
			tmp = make(map[byte]int)
			tmp[s[left]] = left
		} else {
			tmp[s[right]] = right
			right++
		}
	}
	return max(res, len(tmp))
}
