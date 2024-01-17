import "sort"

func minimumCost(nums []int) int64 {
    n, mid := len(nums), 0
    sort.Ints(nums)
    if n%2 == 1 {
        mid = nums[n/2]
    } else {
        mid = (nums[n/2-1] + nums[n/2]) / 2
    }
    left, right := mid, mid
    for ; !isPalindromic(left); left-- {
    }
    for ; !isPalindromic(right); right++ {
    }
    return min(absSum(nums, left), absSum(nums, right))
}

func isPalindromic(num int) bool {
    reverse := 0
    for tmp := num; tmp > 0; tmp /= 10 {
        reverse = 10*reverse + tmp%10
    }
    return reverse == num
}

func absSum(nums []int, x int) int64 {
    sum := int64(0)
    for _, num := range nums {
        sum += int64(abs(x - num))
    }
    return sum
}

func abs(num int) int {
    if num < 0 {
        return -num
    }
    return num
}

func min(x, y int64) int64 {
    if x < y {
        return x
    }
    return y
}
