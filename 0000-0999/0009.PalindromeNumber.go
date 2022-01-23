func isPalindrome(x int) bool {
    if x < 0 || (x != 0 && x%10 == 0) {
        return false
    }
    y := 0
    for y < x {
        y = 10*y + x%10
        x /= 10
    }
    return y == x || y/10 == x
    /*
    if x < 0 {
        return false
    }
    nums := []int{}
    for x > 0 {
        nums = append(nums, x%10)
        x /= 10
    }
    for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
        if nums[i] != nums[j] {
            return false
        }
    }
    return true
    */
}