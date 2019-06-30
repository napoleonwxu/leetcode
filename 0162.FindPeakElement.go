// Same with 852. Peak Index in a Mountain Array
func findPeakElement(nums []int) int {
    peak, right := 0, len(nums)-1
    for peak < right {
        mid := (peak+right) >> 1
        if nums[mid] < nums[mid+1] {
            peak = mid + 1
        } else {
            right = mid
        }
    }
    return peak
}