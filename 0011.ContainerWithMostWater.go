func maxArea(height []int) int {
    area := 0
    i, j := 0, len(height)-1
    for i < j {
        if height[i] < height[j] {
            area = max(area, (j-i)*height[i])
            i++
        } else {
            area = max(area, (j-i)*height[j])
            j--
        }
    }
    return area
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}