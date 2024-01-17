func findIndices(nums []int, indexDifference int, valueDifference int) []int {
    minIdx, maxIdx := 0, 0
    for i := indexDifference; i < len(nums); i++ {
        if nums[i-indexDifference] < nums[minIdx] {
            minIdx = i - indexDifference
        }
        if nums[i-indexDifference] > nums[maxIdx] {
            maxIdx = i - indexDifference
        }
        if nums[i]-nums[minIdx] >= valueDifference {
            return []int{minIdx, i}
        }
        if nums[maxIdx]-nums[i] >= valueDifference {
            return []int{maxIdx, i}
        }
    }
    return []int{-1, -1}
}