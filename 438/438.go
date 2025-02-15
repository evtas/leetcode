package main

// 超时写法
func findAnagrams1(s string, p string) []int {
	k := len(p)
	if len(s) < len(p) {
		return []int{}
	}
	res := make([]int, 0)

	for i, j := 0, k-1; j < len(s); i, j = i+1, j+1 {
		flag := true
		tmp := make(map[byte]int)
		for i := 0; i < len(p); i++ {
			tmp[p[i]] += 1
		}
		for l := i; l <= j; l++ {
			if times, ok := tmp[s[l]]; !ok || times <= 0 {
				flag = false
				break
			} else {
				tmp[s[l]] -= 1
			}
		}
		if flag {
			res = append(res, i)
		}
	}
	return res
}

func findAnagrams(s string, p string) []int {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return []int{}
	}
	res := make([]int, 0)
	sCount, pCount := [26]int{}, [26]int{}
	for i, ch := range p {
		sCount[s[i]-'a'] += 1
		pCount[ch-'a'] += 1
	}
	if sCount == pCount {
		res = append(res, 0)
	}
	for i, ch := range s[:sLen-pLen] {
		sCount[ch-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			res = append(res, i+1)
		}
	}
	return res
}
