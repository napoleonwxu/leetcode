func findLongestWord(s string, d []string) string {
    ans := ""
    for _, w := range d {
        if (len(w) > len(ans) || (len(w) == len(ans) && w < ans)) && isSub(s, w) {
            ans = w
        }
    }
    return ans
}

func isSub(from string, to string) bool {
    n, i := len(to), 0
    for j := 0; j < len(from) && i < n; j++ {
        if from[j] == to[i] {
            i++
        }
    }
    return i == n
}