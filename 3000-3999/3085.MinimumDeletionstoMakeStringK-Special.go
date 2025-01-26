func minimumDeletions(word string, k int) int {
    cnts := make([]int, 26)
    for _, ch := range word {
        cnts[ch-'a']++
    }
    sort.Ints(cnts)
    ans, deleted := len(word), 0
    for i := range cnts {
        cur := deleted
        for j := len(cnts)-1; j > i && cnts[j]-cnts[i] > k; j-- {
            cur += cnts[j] - cnts[i] - k
        }
        ans = min(ans, cur)
        deleted += cnts[i]
    }
    return ans
}