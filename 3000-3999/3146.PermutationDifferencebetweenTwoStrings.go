func findPermutationDifference(s string, t string) int {
    m := make(map[byte]int, len(s))
    for i := range s {
        m[s[i]] = i
    }
    ans := 0
    for i := range t {
        ans += abs(i-m[t[i]])
    }
    return ans
}

func abs(num int) int {
    if num < 0 {
        return -num
    }
    return num
}