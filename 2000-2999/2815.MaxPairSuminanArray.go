func maxSum(nums []int) int {
    ans := -1
    m := make([]int, 10)
    for _, num := range nums {
        d := maxDigit(num)
        if m[d] > 0 {
            ans = max(ans, m[d]+num)
        }
        m[d] = max(m[d], num)
    }
    return ans
}

func maxDigit(num int) int {
    ans := 0
    for ; num > 0; num /= 10 {
        ans = max(ans, num%10)
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}