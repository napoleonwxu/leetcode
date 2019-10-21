func removeDuplicates(nums []int) int {
    i := 0
    for _, num := range nums {
        if i == 0 || num > nums[i-1] {
            nums[i] = num
            i++
        }
    }
    /*
    for j := 0; j < len(nums); j++ {
        if i == 0 || nums[j] != nums[i-1] {
            nums[i] = nums[j]
            i++
        }
    }*/
    return i
}