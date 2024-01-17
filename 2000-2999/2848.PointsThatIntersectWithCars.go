func numberOfPoints(nums [][]int) int {
    points := make([]int, 102)
    for _, num := range nums {
        points[num[0]]++
        points[num[1]+1]--
    }
    for i := 1; i < len(points); i++ {
        points[i] += points[i-1]
    }
    cnt := 0
    for _, point := range points {
        if point > 0 {
            cnt++
        }
    }
    return cnt
}
