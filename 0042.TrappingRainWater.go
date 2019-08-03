func trap(height []int) int {
    n := len(height)
    if n < 3 {
        return 0
    }
    ans := 0
    // Two pointers, O(n) + O(1)
    l, r := 0, n-1
    maxL, maxR := 0, 0
    for l < r {
        if height[l] < height[r] {
            maxL = max(maxL, height[l])
            ans += maxL - height[l]
            l++
        } else {
            maxR = max(maxR, height[r])
            ans += maxR - height[r]
            r--
        }
    }
    /* DP, O(3n) + O(2n)
    maxL, max_tmp := make([]int, n), 0
    for i := 0; i < n; i++ {
        max_tmp = max(max_tmp, height[i])
        maxL[i] = max_tmp
    }
    maxR, max_tmp := make([]int, n), 0
    for i := n-1; i >= 0; i-- {
        max_tmp = max(max_tmp, height[i])
        maxR[i] = max_tmp
    }
    for i, h := range height {
        ans += min(maxL[i], maxR[i]) - h
    } */
    return ans
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