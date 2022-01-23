func searchRange(nums []int, target int) []int {
    ans := []int{-1, -1}
    if len(nums) == 0 {
        return ans
    }
    left, right := 0, len(nums)-1
    for left < right {
        mid := (left+right) >> 1
        if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    if nums[left] != target {
        return ans
    }
    ans[0] = left
    right = len(nums)-1
    for left < right {
        mid := (left+right) >> 1 + 1
        if nums[mid] > target {
            right = mid - 1
        } else {
            left = mid
        }
    }
    ans[1] = right
    return ans
}

func searchRange(nums []int, target int) []int {
    n := len(nums)
    left, right := 0, n-1
    for left <= right {
        mid := (left+right) >> 1
        if nums[mid] == target {
            i, j := mid, mid
            for ; i>0 && nums[i-1]==nums[mid]; i-- {}
            for ; j<n-1 && nums[j+1]==nums[mid]; j++ {}
            return []int{i, j}
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return []int{-1, -1}
}