func longestNiceSubarray(nums []int) int {
    pre, ans := 0, 0
    and := 0
    for i, num := range nums {
        for and&num != 0 {
            and ^= nums[pre]
            pre++
        }
        and |= num
        ans = max(ans, i-pre+1)
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}