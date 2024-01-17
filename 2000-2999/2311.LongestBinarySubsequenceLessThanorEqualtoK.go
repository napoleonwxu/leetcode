func longestSubsequence(s string, k int) int {
    cnt := 0
    num, base := 0, 1
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '0' || num+base <= k {
            num += int(s[i]-'0') * base
            cnt++
        }
        if base <= k { // in case base overflow
            base <<= 1
        }
    }
    return cnt
}