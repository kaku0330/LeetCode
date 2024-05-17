package besttimetobuyandsellstock

func MaxProfit(prices []int) int {
	// day := map[string]int{
	// 	"min":100000,
	// 	"max":0,
	// }
	minIndex, maxIndex := 0, 0
	for i := 0; i < len(prices); i++ {
		if prices[minIndex] > prices[i] {
			minIndex = i
		}
		if prices[maxIndex] < prices[i] {
			maxIndex = i
		}
	}

	if maxIndex > minIndex {
		return 0
	}
	return prices[maxIndex] - prices[minIndex]
}
