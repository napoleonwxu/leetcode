func minSetSize(arr []int) int {
    Map := make(map[int]int)
    for _, num := range arr {
        Map[num]++
    }
    cnt := make([]int, len(Map))
    sum, i := 0, 0
    for _, c := range Map {
        cnt[i] = c
        sum += c
        i++
    }
    sort.Ints(cnt)
    tmp, j := 0, len(cnt)-1
    for tmp < sum/2 {
        tmp += cnt[j]
        j--
    }
    return len(cnt)-j-1
}