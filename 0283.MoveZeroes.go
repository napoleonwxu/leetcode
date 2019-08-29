func moveZeroes(nums []int)  {
    for i, j := 0, 0; j < len(nums); j++ {
        if nums[j] != 0 {
            if i != j {
                nums[i], nums[j] = nums[j], nums[i]
            }
            i++
        }
    }
}