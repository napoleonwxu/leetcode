func leftRigthDifference(nums []int) []int {
    leftSum, rightSum := 0, 0
    for _, num := range nums {
        rightSum += num
    }
    ans := make([]int, len(nums))
    for i, num := range nums {
        rightSum -= num
        ans[i] = abs(leftSum - rightSum)
        leftSum += num
    }
    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}