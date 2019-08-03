func numSubarrayBoundedMax(A []int, L int, R int) int {
    ans := 0
    idx, cnt := -1, 0
    for i, a := range A {
        if a > R {
            idx = i
            cnt = 0
        } else if a >= L {
            cnt = i - idx
        }
        ans += cnt
    }
    return ans
}