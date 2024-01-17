func findArray(pref []int) []int {
    n := len(pref)
    ans := make([]int, n)
    ans[0] = pref[0]
    for i := 1; i < n; i++ {
        ans[i] = pref[i] ^ pref[i-1]
    }
    return ans
}