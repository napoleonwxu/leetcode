func findMaxK(nums []int) int {
    m := make(map[int]bool, len(nums))
    for _, num := range nums {
        m[num] = true
    }
    ans := -1
    for _, num := range nums {
        if num > 0 && m[-num] && num > ans {
            ans = num
        }
    }
    return ans
}