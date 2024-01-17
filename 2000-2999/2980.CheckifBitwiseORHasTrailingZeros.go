func hasTrailingZeros(nums []int) bool {
    even := 0
    for _, num := range nums {
        if num&1 == 0 {
            even++
            if even >= 2 {
                return true
            }
        }
    }
    return false
}