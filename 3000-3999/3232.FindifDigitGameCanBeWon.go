func canAliceWin(nums []int) bool {
    sum1, sum2 := 0, 0
    for _, num := range nums {
        if num < 10 {
            sum1 += num
        } else if num < 100 {
            sum2 += num
        }
    }
    return sum1 != sum2
}