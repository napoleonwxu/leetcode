import "sort"

func minimizeMax(nums []int, p int) int {
    sort.Ints(nums)
    n := len(nums)
    left, right := 0, nums[n-1]-nums[0]
    for left < right {
        mid := (left + right) / 2
        k := 0
        for i := 1; i < n && k < p; i++ {
            if nums[i]-nums[i-1] <= mid {
                k++
                i++
            }
        }
        if k >= p {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}