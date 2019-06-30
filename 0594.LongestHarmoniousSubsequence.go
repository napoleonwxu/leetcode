func findLHS(nums []int) int {
    Map := make(map[int]int)
    cnt := 0
    for _, n := range nums {
        Map[n]++
    }
    for k, v := range Map {
        if Map[k+1] > 0 && v+Map[k+1] > cnt{
            cnt = v+Map[k+1]
        }
    }
    return cnt
}