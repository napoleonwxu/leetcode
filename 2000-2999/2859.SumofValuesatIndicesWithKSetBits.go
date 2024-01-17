func sumIndicesWithKSetBits(nums []int, k int) int {
    sum := 0
    for i, num := range nums {
        ones := 0
        for ; i > 0 && ones <= k; ones++ {
            i &= i - 1
        }
        if ones == k {
            sum += num
        }
    }
    return sum
}