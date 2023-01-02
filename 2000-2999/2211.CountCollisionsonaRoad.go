func countCollisions(directions string) int {
    left, right := 0, len(directions)-1
    for ; left <= right && directions[left] == 'L'; left++ {}
    for ; left <= right && directions[right] == 'R'; right-- {}
    cnt := 0
    for i := left; i <= right; i++ {
        if directions[i] != 'S' {
            cnt++
        }
    }
    return cnt
}
