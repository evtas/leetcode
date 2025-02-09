package main

import "sort"

// 排序+哈希
// 还有第二种方法，数组也能做哈希键？！
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		bs := []byte(str)
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		ss := string(bs)
		println(ss)
		if _, ok := m[ss]; !ok {
			m[ss] = make([]string, 0)
		}
		m[ss] = append(m[ss], str)
	}
	res := make([][]string, 0, len(m))
	for _, val := range m {
		res = append(res, val)
	}
	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groupAnagrams(strs)
}
