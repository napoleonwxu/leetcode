func minAreaRect(points [][]int) int {
    n := len(points)
    Map := make(map[[2]int]bool, n)
    for _, p := range points {
        Map[[2]int{p[0], p[1]}] = true
    }
    ans := math.MaxInt32
    for i := 0; i < n-1; i++ {
        for j := i+1; j < n; j++ {
            p1 := [2]int{points[i][0], points[j][1]}
            p2 := [2]int{points[j][0], points[i][1]}
            tmp := abs((p1[0]-p2[0])*(p1[1]-p2[1]))
            if tmp != 0 && tmp < ans && Map[p1] && Map[p2] {
                ans = tmp
            }
        }
    }
    if ans == math.MaxInt32 {
        return 0
    }
    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
