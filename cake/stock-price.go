package cake

// Writing programming interview questions hasn't made me rich yet ... so I might give up and start trading Apple stocks all day instead.
// https://www.interviewcake.com/question/python/stock-price

// compare every pair
func stockPriceBruteForce(stocks []int) int {
	var maxprof int = -100

	for outer, _ := range stocks {
		for inner, _ := range stocks {
			if outer == inner {
				continue
			}

			earlier := min(outer, inner)
			later := max(outer, inner)

			earlierPrice := stocks[earlier]
			laterPrice := stocks[later]

			profit := laterPrice - earlierPrice
			maxprof = max(maxprof, profit)
		}
	}

	return maxprof
}

// improve the brute force, by achieving O(n) time and O(1) space
func stockPriceBruteForce(stocks []int) int {
	var maxprof int = -100

	for outer, _ := range stocks {
		for inner, _ := range stocks {
			if outer == inner {
				continue
			}

			earlier := min(outer, inner)
			later := max(outer, inner)

			earlierPrice := stocks[earlier]
			laterPrice := stocks[later]

			profit := laterPrice - earlierPrice
			maxprof = max(maxprof, profit)
		}
	}

	return maxprof
}
func min(is ...int) int {
	min := is[0]
	for _, i := range is[1:] {
		if i < min {
			min = i
		}
	}
	return min
}

func max(is ...int) int {
	max := is[0]
	for _, i := range is[1:] {
		if i > max {
			max = i
		}
	}
	return max
}

func StockPrice(stocks []int) int {
	return stockPriceBruteForce(stocks)
}
