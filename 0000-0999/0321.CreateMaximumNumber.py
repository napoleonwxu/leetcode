class Solution(object):
    def maxNumber(self, nums1, nums2, k):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :type k: int
        :rtype: List[int]
        """
        ans = []
        for i in xrange(k+1):
            if i <= len(nums1) and k-i <= len(nums2):
                ans = max(ans, self.merge(self.prep(nums1, i), self.prep(nums2, k-i)))
        return ans

    def prep(self, nums, n):
        ans = []
        drop = len(nums) - n
        for num in nums:
            while ans and drop and ans[-1] < num:
                ans.pop()
                drop -= 1
            ans.append(num)
        return ans[:n]

    def merge(self, nums1, nums2):
        return [max(nums1, nums2).pop(0) for _ in nums1+nums2]
        329. Longest Increasing Path in a Matrix