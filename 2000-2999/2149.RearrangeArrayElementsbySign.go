func rearrangeArray(nums []int) []int {
    ans := make([]int, len(nums))
    posIdx, negIdx := 0, 1
    for _, num := range nums {
        if num > 0 {
            ans[posIdx] = num
            posIdx += 2
        } else {
            ans[negIdx] = num
            negIdx += 2
        }
    }
    return ans
}