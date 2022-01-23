//https://leetcode.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/discuss/531840/JavaC%2B%2BPython-One-Pass

func findTheLongestSubstring(s string) int {
    Map := make(map[int]int)
    Map[0] = -1
    cur, ans := 0, 0
    for i, ch := range s {
        cur ^= 1 << (strings.Index("aeiou", string(ch)) + 1) >> 1
        if _, ok := Map[cur]; !ok {
            Map[cur] = i
        }
        ans = max(ans, i-Map[cur])
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}