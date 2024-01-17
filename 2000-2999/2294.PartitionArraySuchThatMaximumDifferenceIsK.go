import "sort"

func partitionArray(nums []int, k int) int {
    sort.Ints(nums)
    cnt, base := 1, 0
    for i := range nums {
        if nums[i]-nums[base] > k {
            cnt++
            base = i
        }
    }
    return cnt
}