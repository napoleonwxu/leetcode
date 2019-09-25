func splitArray(nums []int, m int) int {
    max, sum := 0, 0
    for _, num := range nums {
        if num > max {
            max = num
        }
        sum += num
    }
    if m == 1 {
        return sum
    }
    left, right := max, sum
    for left <= right {
        mid := (left+right) >> 1
        if canSplit(nums, m, mid) {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return left
}

func canSplit(nums []int, m, target int) bool {
    sum, cnt := 0, 1
    for _, num := range nums {
        sum += num
        if sum > target {
            sum = num
            cnt++
            if cnt > m {
                return false
            }
        }
    }
    return true
}
