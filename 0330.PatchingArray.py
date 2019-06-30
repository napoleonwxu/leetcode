class Solution(object):
    def minPatches(self, nums, n):
        """
        :type nums: List[int]
        :type n: int
        :rtype: int
        """
        upper = 1
        miss = 0
        i = 0
        while upper <= n:
            if i < len(nums) and nums[i] <= upper:
                upper += nums[i]
                i += 1
            else:
                upper += upper
                miss += 1
        return miss

print Solution().minPatches([1, 5, 10], 20)
print Solution().minPatches([1, 2, 2], 5)