func groupAnagrams(strs []string) [][]string {
    // O(NK)
    Map := make(map[[26]int][]string)
    for _, str := range strs {
        cnt := [26]int{}
        for _, ch := range str {
            cnt[ch-'a']++
        }
        Map[cnt] = append(Map[cnt], str)
    }
    ans := make([][]string, 0, len(Map))
    for _, s := range Map {
        ans = append(ans, s)
    }
    return ans
}