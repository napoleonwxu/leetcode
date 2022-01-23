// max[a/b/c/d/e/f...] => min[b/c/d/e/f...], 
// due to i>1, ans: a/(b/c/d/e/f...)
func optimalDivision(nums []int) string {
    n := len(nums)
    strs := make([]string, n)
    for i, num := range nums {
        strs[i] = strconv.Itoa(num)
    }
    if n == 1 {
        return strs[0]
    }
    if n == 2 {
        return strings.Join(strs, "/")
    }
    return strs[0] + "/(" + strings.Join(strs[1:], "/") + ")"
}