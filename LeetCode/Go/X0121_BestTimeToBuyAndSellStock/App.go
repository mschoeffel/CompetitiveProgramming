package X0121_BestTimeToBuyAndSellStock

import (
	"fmt"
	"strconv"
)

func Main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	result := maxProfit(prices)
	fmt.Println("Result Best Time to Buy and Sell Stock: " + strconv.Itoa(result))

}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	max := 0
	min := prices[0]
	for _, p := range prices {
		profit := p - min
		if profit > max {
			max = profit
		}
		if p < min {
			min = p
		}
	}
	return max
}
