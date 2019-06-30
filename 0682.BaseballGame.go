func calPoints(ops []string) int {
    points := []int{}
    for _, op := range ops {
        n := len(points)
        switch op {
            case "+":
                points = append(points, points[n-1]+points[n-2])
            case "D":
                points = append(points, points[n-1]<<1)
            case "C":
                points = points[:n-1]
            default:
                cur, _ := strconv.Atoi(op)
                points = append(points, cur)
        }
    }
    sum := 0
    for _, p := range points {
        sum += p
    }
    return sum
}