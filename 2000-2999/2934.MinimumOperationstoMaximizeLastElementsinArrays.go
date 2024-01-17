func minOperations(nums1 []int, nums2 []int) int {
    n := len(nums1)
    dp1, dp2 := 0, 0
    mi, ma := min(nums1[n-1], nums2[n-1]), max(nums1[n-1], nums2[n-1])
    for i := 0; i < n; i++ {
        if max(nums1[i], nums2[i]) > ma || min(nums1[i], nums2[i]) > mi {
            return -1
        }
        if nums1[i] > nums1[n-1] || nums2[i] > nums2[n-1] {
            dp1++
        }
        if nums1[i] > nums2[n-1] || nums2[i] > nums1[n-1] {
            dp2++
        }
    }
    return min(dp1, dp2)
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}