func numEquivDominoPairs(dominoes [][]int) int {
    Map := make(map[int]int)
    cnt := 0
    for _, d := range dominoes {
        key1, key2 := 10*d[0]+d[1], 10*d[1]+d[0]
        cnt += Map[key1]
        if key1 != key2 { 
            cnt += Map[key2]
        }
        Map[key1]++
    }
    return cnt
}