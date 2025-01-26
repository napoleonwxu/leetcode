func countPairs(nums []int) int {
    ans, n := 0, len(nums)
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            if equal(nums[i], nums[j]) {
                ans++
            }
        }
    }
    return ans
}

func equal(num1, num2 int) bool {
    if num1 < num2 {
        num1, num2 = num2, num1
    }
    d11, d12 := -1, -1
    d21, d22 := -1, -1
    for num1 > 0 {
        d1, d2 := num1%10, num2%10
        if d1 != d2 {
            if d11 == -1 {
                d11, d21 = d1, d2
            } else if d12 == -1 {
                d12, d22 = d1, d2
            } else {
                return false
            }
        }
        num1 /= 10
        num2 /= 10
    }
    return d11 == d22 && d12 == d21
}