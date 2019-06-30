class Solution(object):
    def rob(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        '''
        n = len(nums) + 2
        dp = [0] * n
        for i in xrange(2, n):
            dp[i] = max(dp[i-1], dp[i-2]+nums[i-2])
        return dp[-1]
        '''
        last, now = 0, 0
        for num in nums:
            last, now = now, max(now, last+num)
        return now