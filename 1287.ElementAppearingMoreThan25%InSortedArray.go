func findSpecialInteger(arr []int) int {
    cnt := make(map[int]int)
    for _, a := range arr {
        cnt[a]++
        if cnt[a] > len(arr)>>2 {
            return a
        }
    }
    return arr[0]
}