func maximumCount(nums []int) int {
    neg := binSearchLeft(nums, 0)
    pos := len(nums) - binSearchRight(nums, 0)
    return max(neg, pos)
}

func binSearchLeft(nums []int, num int) int {
    left, right := 0, len(nums)
    for left < right {
        mid := (left + right) / 2
        if nums[mid] < num {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}

func binSearchRight(nums []int, num int) int {
    left, right := 0, len(nums)
    for left < right {
        mid := (left + right) / 2
        if nums[mid] <= num {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
