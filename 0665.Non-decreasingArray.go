func checkPossibility(nums []int) bool {
    n := len(nums)
    for i := 0; i < n-1; i++ {
        if nums[i] > nums[i+1] {
            nums1 := make([]int, n)
            copy(nums1, nums)
            nums1[i] = nums1[i+1]
            nums2 := make([]int, n)
            copy(nums2, nums)
            nums2[i+1] = nums2[i]
            return sort.IntsAreSorted(nums1) || sort.IntsAreSorted(nums2)
            //return sort.IsSorted(sort.IntSlice(nums1)) || sort.IsSorted(sort.IntSlice(nums2))
        }
    }
    return true
    /*
    p := -1
    for i := 0; i < n-1; i++ {
        if nums[i] > nums[i+1] {
            if p >= 0 {
                // exists two nums[p] > nums[p+1]
                return false
            }
            p = i
        }
    }
    // p==-1: no nums[p] > nums[p+1]
    if p <= 0 || p == n-2 {
        return true
    }
    if nums[p-1] <= nums[p+1] || nums[p] <= nums[p+2] {
        return true
    }
    return false
    */
}