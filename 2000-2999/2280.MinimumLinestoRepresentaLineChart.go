import "sort"

func minimumLines(stockPrices [][]int) int {
    if len(stockPrices) < 2 {
        return 0
    }
    sort.Slice(stockPrices, func(i, j int) bool {
        return stockPrices[i][0] < stockPrices[j][0]
    })
    cnt := 1
    for i := 1; i < len(stockPrices)-1; i++ {
        if (stockPrices[i][1]-stockPrices[i-1][1])*(stockPrices[i+1][0]-stockPrices[i][0]) !=
            (stockPrices[i][0]-stockPrices[i-1][0])*(stockPrices[i+1][1]-stockPrices[i][1]) {
            cnt++
        }
    }
    return cnt
}