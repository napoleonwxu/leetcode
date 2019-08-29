func subsets(nums []int) [][]int {
    p := 1 << uint(len(nums))
    ans := make([][]int, p)
    for i := 0; i < p; i++ {
        sub := []int{}
        for j := 0; j < len(nums); j++ {
            if i&(1<<uint(j)) != 0 {
                sub = append(sub, nums[j])
            }
        }
        ans[i] = sub
    }
    return ans
}