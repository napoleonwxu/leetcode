func pivotIndex(nums []int) int {
    sum, suml := 0, 0
    for _, v := range nums {
        sum += v
    }
    for i, v := range nums {
        if suml == sum - suml - v {
            return i
        }
        suml += v
    }
    return -1
}