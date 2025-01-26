func maxRectangleArea(points [][]int) int {
    ans, n := -1, len(points)
    sort.Slice(points, func(i, j int) bool {
        if points[i][0] == points[j][0] {
            return points[i][1] < points[j][1]
        }
        return points[i][0] < points[j][0]
    })
    for i := 0; i < n-3; i++ {
        if points[i][0] != points[i+1][0] {
            continue
        }
        for j := i+2; j < n-1; j++ {
            if points[j][1] >= points[i][1] && points[j][1] <= points[i+1][1] {
                ans = max(ans, getArea(points, i, i+1, j, j+1))
                break
            }
        }
    }
    return ans
}

func getArea(p [][]int, i, j, k, l int) int {
    if p[i][0] == p[j][0] && p[k][0] == p[l][0] && 
        p[i][1] == p[k][1] && p[j][1] == p[l][1] {
            return (p[k][0]-p[i][0]) * (p[j][1]-p[i][1])
        }
    return -1
}