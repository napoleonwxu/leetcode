func sortArray(nums []int) []int {
    qSort(nums, 0, len(nums)-1)
    return nums
}

func qSort(nums []int, left, right int) {
    if left >= right {
        return
    }
    mid := partition(nums, left, right)
    qSort(nums, left, mid-1)
    qSort(nums, mid+1, right)
}

func partition(nums []int, left, right int) int {
    mid := left
    for i := left; i < right; i++ {
        if nums[i] < nums[right] {
            nums[i], nums[mid] = nums[mid], nums[i]
            mid++
        }
    }
    nums[mid], nums[right] = nums[right], nums[mid]
    return mid
}