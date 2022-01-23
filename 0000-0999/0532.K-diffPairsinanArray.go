func findPairs(nums []int, k int) int {
    if k < 0 {
        return 0
    }
    Map := make(map[int]int)
    for _, num := range nums {
        Map[num]++
    }
    cnt := 0
    if k == 0 {
        for _, c := range Map {
            if c >= 2 {
                cnt++
            }
        }
        return cnt
    }
    for num, _ := range Map {
        if Map[num+k] > 0 {
            cnt++
        }
    }
    return cnt
}