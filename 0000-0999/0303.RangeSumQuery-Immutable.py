class NumArray(object):
    def __init__(self, nums):
        """
        initialize your data structure here.
        :type nums: List[int]
        """
        self.dp = nums
        for i in xrange(1, len(nums)):
            self.dp[i] += self.dp[i-1]

    def sumRange(self, i, j):
        """
        sum of elements nums[i..j], inclusive.
        :type i: int
        :type j: int
        :rtype: int
        """
        '''
        if i == 0:
            return self.dp[j]
        else:
            return self.dp[j] - self.dp[i-1]
        '''
        return self.dp[j] - (self.dp[i-1] if i else 0)


# Your NumArray object will be instantiated and called as such:
# numArray = NumArray(nums)
# numArray.sumRange(0, 1)
# numArray.sumRange(1, 2)