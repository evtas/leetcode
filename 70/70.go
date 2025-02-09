package _0

var valMap = map[int]int{
	1: 1,
	2: 2,
}

func climbStairs(n int) int {
	if val, ok := valMap[n]; ok {
		return val
	}
	res := climbStairs(n-1) + climbStairs(n-2)
	valMap[n] = res
	return res
}

func climbStairs2(n int) int {
	p, q, r := 1, 1, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
