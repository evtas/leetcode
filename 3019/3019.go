package main

import "strings"

func countKeyChanges(s string) int {
	if len(s) <= 1 {
		return 0
	}
	res := 0
	sLower := strings.ToLower(s)
	for i := 1; i < len(sLower); i++ {
		if sLower[i] != sLower[i-1] {
			res += 1
		}
	}
	return res
}

func main() {
	res := countKeyChanges("AaAaAaaA")
	println(res)
}
