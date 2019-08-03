func dominantIndex(nums []int) int {
    // One Pass
    max, sec := math.MinInt32, math.MinInt32
    maxIdx := -1
    for i, num := range nums {
        if num > max {
            max, sec = num, max
            maxIdx = i
        //} else if num > sec {
        } else if num != max && num > sec { // In case max is not unique
            sec = num
        }
    }
    if 2*sec > max {
        return -1
    }
    /* O(2n)
    maxIdx := 0
    for i, num := range nums {
        if num > nums[maxIdx] {
            maxIdx = i
        }
    }
    for i, num := range nums {
        if i != maxIdx && num > nums[maxIdx]>>1 {
            return -1
        }
    } */
    return maxIdx
}