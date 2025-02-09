package _09

func fib(n int) int {
	p, q, r := 0, 1, 0
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
