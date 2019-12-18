func splitArray(nums []int, m int) int {
    left, right := 0, 0
    for _, num := range nums {
        left = max(left, num)
        right += num
    }
    for left < right {
        mid := (left+right) >> 1
        if split(nums, mid) <= m {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}

func split(nums []int, target int) int {
    cnt, sum := 1, 0
    for _, num := range nums {
        sum += num
        if sum > target {
            cnt++
            sum = num
        }
    }
    return cnt
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}