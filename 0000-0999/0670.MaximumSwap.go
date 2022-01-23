func maximumSwap(num int) int {
    // O(n)
    digits := []int{}
    tmp := num
    for tmp > 0 {
        digits = append([]int{tmp%10}, digits...)
        tmp /= 10
    }
    idx := make(map[int]int, 10)
    for i, d := range digits {
        idx[d] = i
    }
    for i, d := range digits {
        for n := 9; n > d; n-- {
            if j, _ := idx[n]; j > i {
                digits[i], digits[j] = digits[j], digits[i]
                return toInt(digits)
            }
        }
    }
    return num
}

func toInt(nums []int) int {
    num := 0
    for _, n := range nums {
        num = 10*num + n
    }
    return num
}