func maxPoints(points [][]int) int {
    n := len(points)
    if n <= 2 {
        return n
    }
    ans := 0
    for i := 0; i < n; i++ {
        Map := make(map[[2]int]int)
        same, curMax := 0, 0
        for j := i+1; j < n; j++ {
            x, y := points[i][0]-points[j][0], points[i][1]-points[j][1]
            if x == 0 && y == 0 {
                same++
                continue
            }
            g := gcd(x, y)
            if g != 0 {
                x, y = x/g, y/g
            }
            Map[[2]int{x,y}]++
            curMax = max(curMax, Map[[2]int{x,y}])
        }
        ans = max(ans, curMax+same+1)
    }
    return ans
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x%y
    }
    return x
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}