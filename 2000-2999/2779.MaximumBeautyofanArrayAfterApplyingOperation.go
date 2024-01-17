import "sort"

func maximumBeauty(nums []int, k int) int {
    sort.Ints(nums)
    l, r := 0, 0
    for ; r < len(nums); r++ {
        if nums[r]-nums[l] > 2*k {
            l++
        }
    }
    return r - l
}