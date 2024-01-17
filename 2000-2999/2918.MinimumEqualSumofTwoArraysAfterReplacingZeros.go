func minSum(nums1 []int, nums2 []int) int64 {
    sum1, zero1 := sum(nums1)
    sum2, zero2 := sum(nums2)
    if sum1 < sum2 && zero1 == 0 {
        return -1
    }
    if sum2 < sum1 && zero2 == 0 {
        return -1
    }
    return max(sum1, sum2)
}

func sum(nums []int) (int64, int) {
    sum, zero := int64(0), 0
    for _, num := range nums {
        sum += max(int64(num), 1)
        if num == 0 {
            zero++
        }
    }
    return sum, zero
}

func max(x, y int64) int64 {
    if x > y {
        return x
    }
    return y
}