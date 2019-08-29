// https://leetcode.com/problems/maximum-of-absolute-value-expression/discuss/339968/JavaC%2B%2BPython-Maximum-Manhattan-Distance
func maxAbsValExpr(arr1 []int, arr2 []int) int {
    n := len(arr1)
    ans := 0
    // O(n)
    dir := [2]int{-1, 1}
    for _, x := range dir {
        for _, y := range dir {
            closest := x*arr1[0] + y*arr2[0] + 0
            for i := 1; i < n; i++ {
                cur := x*arr1[i] + y*arr2[i] + i
                ans = max(ans, cur-closest)
                closest = min(closest, cur)
            }
        }
    }
    /* O(n^2)
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            ans = max(ans, abs(arr1[i]-arr1[j]) + abs(arr2[i]-arr2[j]) + abs(i-j))
        }
    } */
    return ans
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}