func minimumPushes(word string) int {
    n := len(word)
    ans, push := 0, 1
    for n > 0 {
        cnt := min(8, n)
        ans += push * cnt
        push++
        n -= cnt
    }
    return ans
}