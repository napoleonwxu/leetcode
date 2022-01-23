func maxArea(height []int) int {
    ans := 0
    left, right := 0, len(height)-1
    for left < right {
        if height[left] <= height[right] {
            ans = max(ans, (right-left) * height[left])
            left++
        } else {
            ans = max(ans, (right-left) * height[right])
            right--
        }
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}