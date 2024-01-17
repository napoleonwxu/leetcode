import "sort"

func divideArray(nums []int, k int) [][]int {
    n := len(nums)
    sort.Ints(nums)
    ans := make([][]int, n/3)
    for i := 2; i < n; i += 3 {
        if nums[i]-nums[i-2] > k {
            return [][]int{}
        }
        ans[i/3] = []int{nums[i-2], nums[i-1], nums[i]}
    }
    return ans
}