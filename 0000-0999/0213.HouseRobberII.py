class Solution(object):
    def rob(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        n = len(nums)
        if n <= 1:
            return sum(nums)
        return max(self.dp(nums, 0, n-1), self.dp(nums, 1, n))

    def dp(self, nums, l, r):
        last, now = 0, 0
        for num in nums[l:r]:
            last, now = now, max(now, last+num)
        return now
