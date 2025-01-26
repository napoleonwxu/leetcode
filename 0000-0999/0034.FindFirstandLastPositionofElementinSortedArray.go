func searchRange(nums []int, target int) []int {
    left := insertRight(nums, target-1)
    right := insertRight(nums, target)
    if left == right {
        return []int{-1, -1}
    }
    return []int{left, right-1}
    /*
    left := insertLeft(nums, target)
    if left < 0 || left >= len(nums) || nums[left] != target {
        return []int{-1, -1}
    }
    right := insertRight(nums, target)
    return []int{left, right-1}
    */
}

func insertLeft(nums []int, target int) int {
    left, right := 0, len(nums)
    for left < right {
        mid := (left + right) / 2
        if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}

func insertRight(nums []int, target int) int {
    left, right := 0, len(nums)
    for left < right {
        mid := (left + right) / 2
        if nums[mid] <= target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
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