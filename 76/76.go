package main

import "math"

func minWindow(s string, t string) string {
	oriM, cntM := make(map[byte]int), make(map[byte]int)

	for i := 0; i < len(t); i++ {
		oriM[t[i]]++
	}

	check := func() bool {
		for k, v := range oriM {
			if cntM[k] < v {
				return false
			}
		}
		return true
	}
	sLen := len(s)
	ll := math.MaxInt32
	ansL, ansR := -1, -1

	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && oriM[s[r]] != 0 {
			cntM[s[r]]++
		}
		for check() && l <= r {
			if (r - l + 1) < ll {
				ll = r - l + 1
				ansL, ansR = l, l+ll
			}
			if oriM[s[l]] != 0 {
				cntM[s[l]]--
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}
