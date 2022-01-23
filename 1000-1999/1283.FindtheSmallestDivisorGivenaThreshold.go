func smallestDivisor(nums []int, threshold int) int {
    left, right := 1, max(nums)
    for left < right {
        mid := (left+right) >> 1
        if divSum(nums, mid) <= threshold {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}

func divSum(nums []int, div int) int {
    sum := 0
    for _, num := range nums {
        sum += (num+div-1)/div
        /*
        sum += num/div
        if num%div != 0 {
            sum++
        }*/
    }
    return sum
}

func max(nums []int) int {
    m := 1
    for _, num := range nums {
        if num > m {
            m = num
        }
    }
    return m
}
