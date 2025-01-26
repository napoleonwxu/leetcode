func maxLength(nums []int) int {
    ans := 0
    for i := 0; i < len(nums); i++ {
        curGCD, curLCM := nums[i], nums[i]
        curProd := nums[i]
        for j := i+1; j < len(nums); j++ {
            curGCD = gcd(curGCD, nums[j])
            curLCM = lcm(curLCM, nums[j])
            curProd *= nums[j]
            if curProd == curGCD*curLCM {
                ans = max(ans, j-i+1)
            }
        }
    }
    return ans
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x%y
    }
    return x
}

func lcm(x, y int) int {
    return x / gcd(x, y) * y
}