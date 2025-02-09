package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := ""
	s := strs[0]
	for i := 0; i < len(s); i++ {
		cur := s[i]
		for j := 1; j < len(strs); j++ {
			if i > len(strs[j])-1 || strs[j][i] != cur {
				return res
			}
		}
		res += string(cur)
	}
	return res
}

func main() {
	res := longestCommonPrefix([]string{"ab", "a"})
	println(res)
}
