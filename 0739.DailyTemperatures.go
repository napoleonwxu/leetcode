func dailyTemperatures(T []int) []int {
    n := len(T)
    ans := make([]int, n)
    // O(n) with stack
    stack := []int{}
    for i := n-1; i >= 0; i-- {
        j := len(stack) - 1
        for ; j >= 0 && T[stack[j]] <= T[i]; j-- {}
        if j >= 0 {
            ans[i] = stack[j] - i
        }
        stack = append(stack[:j+1], i)
    }
    /* O(n^2)
    for i, t := range T {
        for j := i+1; j < n; j++ {
            if T[j] > t {
                ans[i] = j - i
                break
            }
        }
    }*/
    return ans
}