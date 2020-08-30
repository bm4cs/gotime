package cake

import "fmt"

// Writing programming interview questions hasn't made me rich yet ... so I might give up and start trading Apple stocks all day instead.
// https://www.interviewcake.com/question/python/stock-price

// compare every pair: O(n^2) time and O(1) space
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
func stockPriceGreedy(stocks []int) int {
	var max_prof int = -100
	var min_price int = stocks[0]

	for i, e := range stocks[1:] {
		fmt.Printf("index %d price %d ", i, e)
		profit := e - min_price
		max_prof = max(max_prof, profit)
		fmt.Printf("profit %d max_prof %d min_price %d\n", profit, max_prof, min_price)

		// update min price before next iteration
		if (e < min_price) {
			min_price = e;
		}
	}

	return max_prof
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

func StockPrice() {
	fmt.Println("Running StockPrice puzzle")
	stocks := []int{10, 7, 5, 8, 11, 9} //6
	//stocks := []int{10, 7, 5, 3, 1} //-2
	//max_profit := stockPriceBruteForce(stocks)
	max_profit := stockPriceGreedy(stocks)
	fmt.Println(max_profit)
}
