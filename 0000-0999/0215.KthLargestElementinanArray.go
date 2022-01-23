// ByteDance-Lark
func findKthLargest(nums []int, k int) int {
    /* quick select: 
    average: O(n) + O(n/2) + O(n/4) + O(n/8) + ... = O(2n) = O(n)
    worst case: O(n^2)
    */
    left, right := 0, len(nums)-1
    for left < right {
        mid := partition(nums, left, right)
        if mid == k-1 {
            break
        } else if mid < k-1 {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return nums[k-1]
}

func partition(nums []int, l, r int) int {
    mid := l
    for i := l; i < r; i++ {
        if nums[i] > nums[r] {
            nums[mid], nums[i] = nums[i], nums[mid]
            mid++
        }
    }
    nums[mid], nums[r] = nums[r], nums[mid]
    return mid
}

/*
func partition(nums []int, l, r int) int {
    pivot := nums[l]
    for l < r {
        for ; l < r && nums[r] <= pivot; r-- {}
        nums[l] = nums[r]
        for ; l < r && nums[l] > pivot; l++ {}
        nums[r] = nums[l]
    }
    nums[l] = pivot
    return l
}
*/