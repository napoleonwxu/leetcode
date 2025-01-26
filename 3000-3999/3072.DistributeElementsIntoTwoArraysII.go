func resultArray(nums []int) []int {
    nums1, nums2 := []int{nums[0]}, []int{nums[1]}
    sort1, sort2 := []int{nums[0]}, []int{nums[1]}
    for i := 2; i < len(nums); i++ {
        idx1 := binSearch(sort1, nums[i])
        idx2 := binSearch(sort2, nums[i])
        if (idx1 > idx2) || (idx1 == idx2 && len(nums1) <= len(nums2)) {
            nums1 = append(nums1, nums[i])
            sort1 = insert(sort1, idx1, nums[i])
        } else {
            nums2 = append(nums2, nums[i])
            sort2 = insert(sort2, idx2, nums[i])
        }
    }
    return append(nums1, nums2...)
}

func binSearch(nums []int, target int) int {
    l, r := 0, len(nums) - 1
    for l <= r {
        mid := (l+r) / 2
        if nums[mid] <= target {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return l
}

func insert(nums []int, idx, target int) []int {
    nums = append(nums, 0)
    copy(nums[idx+1:], nums[idx:])
    nums[idx] = target
    return nums
}