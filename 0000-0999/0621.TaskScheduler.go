func leastInterval(tasks []byte, n int) int {
    if n == 0 {
        return len(tasks)
    }
    cnt := make([]int, 26)
    for _, task := range tasks {
        cnt[task-'A']++
    }
    sort.Ints(cnt)
    // O(n)
    cycle := cnt[25] - 1
    extra := 0
    for i := 25; i >= 0 && cnt[i] == cnt[25]; i-- {
        extra++
    }
    return max(len(tasks), cycle*(n+1) + extra)
    /* O(interval)
    interval := 0
    for cnt[25] > 0 {
        i := 0
        for i <= n {
            if cnt[25] == 0 {
                break
            }
            if i < 26 && cnt[25-i] > 0 {
                cnt[25-i]--
            }
            interval++
            i++
        }
        sort.Ints(cnt)
    }
    return interval
    */
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}