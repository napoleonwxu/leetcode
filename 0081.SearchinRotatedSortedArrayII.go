func search(nums []int, target int) bool {
    // average O(logN), worst O(N)
    left, right := 0, len(nums)-1
    for left <= right {
        mid := (left+right) >> 1
        if nums[mid] == target {
            return true
        }
        if nums[mid] == nums[left] && nums[mid] == nums[right] {
            left, right = left+1, right-1
        } else if nums[mid] >= nums[left] {
            if target >= nums[left] && target < nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else {
            if target > nums[mid] && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    return false
}