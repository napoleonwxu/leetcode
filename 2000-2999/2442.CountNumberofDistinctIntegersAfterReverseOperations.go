func countDistinctIntegers(nums []int) int {
    m := make(map[int]bool, 2*len(nums))
    for _, num := range nums {
        m[num] = true
        reverse := 0
        for ; num > 0; num /= 10 {
            reverse = 10*reverse + num%10
        }
        m[reverse] = true
    }
    return len(m)
}