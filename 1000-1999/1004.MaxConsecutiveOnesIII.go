func longestOnes(nums []int, k int) int {
    start, end := 0, 0
    for ; end < len(nums); end++ {
        if nums[end] == 0 {
            k--
        }
        if k < 0 {
            if nums[start] == 0 {
                k++
            }
            start++
        }
    }
    return end-start
}
