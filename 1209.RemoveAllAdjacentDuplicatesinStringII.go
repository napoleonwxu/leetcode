func removeDuplicates(s string, k int) string {
    // O(n)
    arr := make([]byte, len(s))
    cnt := make([]int, len(s))
    n := 0
    for i := 0; i < len(s); i++ {
        if n > 0 && s[i] == arr[n-1] {
            cnt[n-1]++
            if cnt[n-1] == k {
                n--
            }
        } else {
            arr[n], cnt[n] = s[i], 1
            n++
        }
    }
    ans := make([]string, n)
    for i := 0; i < n; i++ {
        ans[i] = strings.Repeat(string(arr[i]), cnt[i])
    }
    return strings.Join(ans, "")
    /* O(n*k)
    stack := make([]byte, len(s))
    n := 0
    for i := 0; i < len(s); i++ {
        stack[n] = s[i]
        n++
        if n >= k && check(stack, n-1, k) {
            n -= k
        }
    }
    return string(stack[:n])
    */
}

func check(stack []byte, idx, k int) bool {
    for i := 1; i < k; i++ {
        if stack[idx-i] != stack[idx] {
            return false
        }
    }
    return true
}