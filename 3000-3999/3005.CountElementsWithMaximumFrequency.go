func maxFrequencyElements(nums []int) int {
    m := make(map[int]int, len(nums))
    for _, num := range nums {
        m[num]++
    }
    ma, ans := 0, 0
    for _, cnt := range m {
        if cnt == ma {
            ans += cnt
        } else if cnt > ma {
            ma = cnt
            ans = cnt
        }
    }
    return ans
}