func numRabbits(answers []int) int {
    Map := make(map[int]int)
    for _, r := range answers {
        Map[r]++
    }
    cnt := 0
    for r, c := range Map {
        //cnt += int(math.Ceil(float64(c)/float64(r+1))) * (r+1)
        cnt += (c+r)/(r+1) * (r+1)
    }
    return cnt
}