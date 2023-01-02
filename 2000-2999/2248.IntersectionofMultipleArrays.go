func intersection(nums [][]int) []int {
    cnt := make(map[int]int)
    for _, num := range nums {
        for _, n := range num {
            cnt[n]++
        }
    }
    ans := make([]int, 0, len(nums[0]))
    for n, c := range cnt {
        if c == len(nums) {
            ans = append(ans, n)
        }
    }
    sort.Ints(ans)
    return ans
}