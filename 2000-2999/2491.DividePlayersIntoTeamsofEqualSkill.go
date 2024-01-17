func dividePlayers(skill []int) int64 {
    sum, n := 0, len(skill)
    m := make(map[int]int, n)
    for _, num := range skill {
        sum += num
        m[num]++
    }
    tmp := 2 * sum / n
    ans := int64(0)
    for num, cnt := range m {
        if cnt != m[tmp-num] {
            return -1
        }
        ans += int64(cnt * num * (tmp - num))
    }
    return ans / 2
}