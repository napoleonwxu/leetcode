// https://leetcode.com/problems/search-in-rotated-sorted-array/discuss/14436/Revised-Binary-Search
func search(nums []int, target int) int {
    //return binSearch(nums, target, 0, len(nums)-1)
    left, right := 0, len(nums)-1
    for left <= right {
        mid := (left+right) >> 1
        if nums[mid] == target {
            return mid
        } else if nums[left] <= nums[mid] {
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
    return -1
}

func binSearch(nums []int, target, l, r int) int {
    if l > r {
        return -1
    }
    mid := (l+r) / 2
    if nums[mid] == target {
        return mid
    }
    if nums[l] <= nums[mid] {
        if nums[l] <= target && target < nums[mid] {
            return binSearch(nums, target, l, mid-1)
        } else {
            return binSearch(nums, target, mid+1, r)
        }
    } else {
        if nums[mid] < target && target <= nums[r] {
            return binSearch(nums, target, mid+1, r)
        } else {
            return binSearch(nums, target, l, mid-1)
        }
    }
}
