func twoSum(numbers []int, target int) []int {
    left, right := 0, len(numbers)-1
    for left < right {
        sum := numbers[left] + numbers[right]
        if sum == target {
            return []int{left+1, right+1}
        } else if sum < target {
            left += 1
        } else {
            right -= 1
        }
    }
    return nil
}