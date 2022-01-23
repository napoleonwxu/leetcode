func findLonely(nums []int) []int {
    m := make(map[int]int)
    for _, num := range nums {
        m[num]++
    }
    ans := make([]int, 0, len(m))
    for num, cnt := range m {
        if cnt == 1 && m[num-1] == 0 && m[num+1] == 0 {
            ans = append(ans, num)
        }
    }
    return ans
}