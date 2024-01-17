import "sort"

func maximumLength(s string) int {
    n := len(s)
    m := make(map[byte][]int, 26)
    for i := 0; i < n; {
        j := i + 1
        for ; j < n && s[j] == s[i]; j++ {
        }
        m[s[i]] = append(m[s[i]], j-i)
        i = j
    }
    ans := -1
    for _, lens := range m {
        sort.Ints(lens)
        cnts := make([]int, lens[len(lens)-1]+1)
        for _, l := range lens {
            for i := 1; i <= l; i++ {
                cnts[i] += l - i + 1
            }
        }
        for c, cnt := range cnts {
            if cnt >= 3 && c > ans {
                ans = c
            }
        }
    }
    return ans
}