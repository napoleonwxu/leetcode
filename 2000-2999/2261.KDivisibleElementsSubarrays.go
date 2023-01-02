// O(n^3)
func countDistinct(nums []int, k int, p int) int {
    n := len(nums)
    m := make(map[string]bool)
    for i := 0; i < n; i++ {
        for j := i+1; j <= n; j++ {
            if h := hash(nums[i:j], k, p); h != "" {
                m[h] = true
            }
        }
    }
    return len(m)
}

func hash(nums []int, k, p int) string {
    strs := make([]string, len(nums))
    for i, num := range nums {
        if num%p == 0 {
            k--
        }
        if k < 0 {
            return ""
        }
        strs[i] = strconv.Itoa(num)
    }
    return strings.Join(strs, ",")
}
