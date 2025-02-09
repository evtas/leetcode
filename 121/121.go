package _21

func maxProfit(prices []int) int {
	minest := prices[0]
	profit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minest {
			minest = prices[i]
		}
		profit = max(profit, prices[i]-minest)
	}
	return profit
}

func main() {

}
