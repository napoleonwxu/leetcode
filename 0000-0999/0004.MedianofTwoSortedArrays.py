class Solution(object):
    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """
        # binary search, O(log(m+n))
        l = len(nums1) + len(nums2)
        if l%2:
            return self.findKth(nums1, nums2, l/2)
        else:
            return (self.findKth(nums1, nums2, l/2-1) + self.findKth(nums1, nums2, l/2)) / 2.0

        ''' # merge, O(m+n) + O(m+n)
        nums = []
        i = j = 0
        while i < len(nums1) and j < len(nums2):
            if nums1[i] <= nums2[j]:
                nums.append(nums1[i])
                i += 1
            else:
                nums.append(nums2[j])
                j += 1
        nums += nums1[i:] + nums2[j:]
        l = len(nums)
        if l % 2:
            return nums[l/2]
        else:
            return (nums[l/2-1]+nums[l/2])/2.0
        '''

    def findKth(self, nums1, nums2, k):
        if len(nums1) > len(nums2):
            nums1, nums2 = nums2, nums1
        if not nums1:
            return nums2[k]
        if len(nums1) + len(nums2) - 1 == k:
            return max(nums1[-1], nums2[-1])
        i = len(nums1) / 2
        j = k - i
        if nums1[i] > nums2[j]:
            return self.findKth(nums1[:i], nums2[j:], i)
        else:
            return self.findKth(nums1[i:], nums2[:j], j)
