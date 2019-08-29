func largestUniqueNumber(A []int) int {
    cnt := [1001]int{}
    for _, a := range A {
        cnt[a]++
    }
    for a := 1000; a >= 0; a-- {
        if cnt[a] == 1 {
            return a
        }
    }
    return -1
}