func largestRectangleArea(heights []int) int {
    heights = append(heights, 0)
    stack := []int{-1}
    ans := 0
    for end, h := range heights {
        for len(stack) > 1 && h < heights[stack[len(stack)-1]] {
            // use stack[len(stack)-2]+1 instead of stack[len(stack)-1] to deal with duplicated values: ex [2,1,2]
            height := heights[stack[len(stack)-1]]
            start := stack[len(stack)-2] + 1
            ans = max(ans, height * (end-start))
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, end)
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}