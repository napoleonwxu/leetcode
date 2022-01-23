func sortColors(nums []int)  {
    l, r := 0, len(nums)-1
    for i := 0; i <= r; {
        if nums[i] == 0 {
            nums[i], nums[l] = nums[l], nums[i]
            l++
            i++
        } else if nums[i] == 2 {
            nums[i], nums[r] = nums[r], nums[i]
            r--
        } else {
            i++
        }
    }
    /*
    for i := 0; i <= r; i++ {
        for i < r && nums[i] == 2 {
            nums[i], nums[r] = nums[r], nums[i]
            r--
        }
        for i > l && nums[i] == 0 {
            nums[i], nums[l] = nums[l], nums[i]
            l++
        }
    }*/
}