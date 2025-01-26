func sumDigitDifferences(nums []int) int64 {
    n, m := len(nums), 0
    for tmp := nums[0]; tmp > 0; tmp /= 10 {
        m++
    }
    cnts := make([][]int, m)
    for i := range cnts {
        cnts[i] = make([]int, 10)
    }
    for _, num := range nums {
        for i := 0; i < m; i++ {
            cnts[i][num%10]++
            num /= 10
        }
    }
    ans := int64(0)
    for _, cnt := range cnts {
        for i := 0; i < len(cnt); i++ {
            ans += int64(cnt[i]*(n-cnt[i]))
        }
    }
    return ans/2
}