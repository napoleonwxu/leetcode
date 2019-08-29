func compress(chars []byte) int {
    n := len(chars)
    left, right := 0, 0
    for right < n {
        cur, cnt := chars[right], 0
        for right < n && chars[right] == cur {
            right++
            cnt++
        }
        chars[left] = cur
        left++
        if cnt > 1 {
            for _, ch := range strconv.Itoa(cnt) {
                chars[left] = byte(ch)
                left++
            }
        }
    }
    return left
}