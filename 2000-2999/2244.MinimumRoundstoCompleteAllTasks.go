func minimumRounds(tasks []int) int {
    cnt := make(map[int]int)
    for _, task := range tasks {
        cnt[task]++
    }
    ans := 0
    for _, c := range cnt {
        if c == 1 {
            return -1
        }
        ans += (c+2) / 3
    }
    return ans
}