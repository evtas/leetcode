package main

func largestGoodInteger(num string) string {
	res := ""
	for i := 2; i < len(num); i++ {
		if num[i] == num[i-1] && num[i] == num[i-2] {
			if res == "" || res[0] < num[i] {
				res = num[i-2 : i+1]
			}
		}
	}
	return res
}

func main() {

}
