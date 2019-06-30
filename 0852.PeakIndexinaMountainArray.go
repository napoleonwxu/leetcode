func peakIndexInMountainArray(A []int) int {
    peak, right := 0, len(A)-1
    for peak < right {
        mid := (peak+right) >> 1
        if A[mid] < A[mid+1] {
            peak = mid + 1
        } else {
            right = mid
        }
    }
    return peak
}