func countLatticePoints(circles [][]int) int {
    MAX := 100 * 2
    cnt := 0
    for i := 0; i <= MAX; i++ {
        for j := 0; j <= MAX; j++ {
            if inCircles(i, j, circles) {
                cnt++
            }
        }
    }
    return cnt
}

func inCircles(i, j int, circles [][]int) bool {
    for _, c := range circles {
        if (i-c[0])*(i-c[0]) + (j-c[1])*(j-c[1]) <= c[2]*c[2] {
            return true
        }
    }
    return false
}