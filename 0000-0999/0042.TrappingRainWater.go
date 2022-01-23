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
    /* DP, O(2n) + O(2n)
    left, right := make([]int, n), make([]int, n)
    left[0], right[n-1] = height[0], height[n-1]
    for i := 1; i < n; i++ {
        left[i] = max(left[i-1], height[i])
        right[n-i-1] = max(right[n-i], height[n-i-1])
    }
    for i := 1; i < n-1; i++ {
        ans += min(left[i], right[i]) - height[i]
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