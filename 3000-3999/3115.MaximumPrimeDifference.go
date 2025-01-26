func maximumPrimeDifference(nums []int) int {
    l, r := 0, len(nums)-1
    for ; l < len(nums); l++ {
        if isPrime(nums[l]) {
            break
        }
    }
    for ; r > 0; r-- {
        if isPrime(nums[r]) {
            break
        }
    }
    return r-l
}

func isPrime(num int) bool {
    if num <= 1 {
        return false
    }
    if num == 2 {
        return true
    }
    if num%2 == 0 {
        return false
    }
    for i := 3; i*i <= num; i += 2 {
        if num%i == 0 {
            return false
        }
    }
    return true
}