func maxRepOpt1(text string) int {
    n := len(text)
    cnt := make([]int, 26)
    for _, ch := range text {
        cnt[int(ch-'a')]++
    }
    ans := 0
    for i := 0; i < n; i++ {
        diff, sum := 0, 0
        for j := i; j < n && (text[j] == text[i] || diff < 1) && sum < cnt[text[i]-'a']; j++ {
            if text[j] != text[i] {
                diff++
            }
            sum++
        }
        if sum > ans {
            ans = sum
        }
    }
    return ans
}