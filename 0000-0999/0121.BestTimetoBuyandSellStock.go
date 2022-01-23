func maxProfit(prices []int) int {
    minPrice, maxProfit := math.MaxInt32, 0
    for _, price := range prices {
        minPrice = min(minPrice, price)
        maxProfit = max(maxProfit, price-minPrice)
    }
    /*
    profit, maxProfit := 0, 0
    for i := 1; i < len(prices); i++ {
        profit += prices[i] - prices[i-1]
        profit = max(0, profit)
        maxProfit = max(maxProfit, profit)
    }*/
    return maxProfit
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}