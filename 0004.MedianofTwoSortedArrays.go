func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    // O(m+n)
    length := len(nums1) + len(nums2)
    if length&1 == 1 {
        return float64(findKth(nums1, nums2, length>>1))
    } else {
        return float64(findKth(nums1, nums2, length>>1-1) + findKth(nums1, nums2, length>>1)) / 2
    }
}

func findKth(nums1, nums2 []int, k int) int {
    if len(nums1) > len(nums2) {
        nums1, nums2 = nums2, nums1
    }
    m, n := len(nums1), len(nums2)
    if m == 0 {
        return nums2[k]
    }
    if m+n-1 == k {
        return max(nums1[m-1], nums2[n-1])
    }
    i := m >> 1
    j := k - i
    if nums1[i] < nums2[j] {
        return findKth(nums1[i:], nums2[:j], j)
    } else {
        return findKth(nums1[:i], nums2[j:], i)
    }
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}