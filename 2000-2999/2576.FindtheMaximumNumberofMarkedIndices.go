import "sort"

func maxNumOfMarkedIndices(nums []int) int {
    sort.Ints(nums)
    n, i := len(nums), 0
    for j := n - n/2; j < n; j++ {
        if 2*nums[i] <= nums[j] {
            i++
        }
    }
    return 2 * i
}