class Solution(object):
    def canJump(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        '''
        if len(nums) == 1:
            return True
        d = 1
        for i in xrange(len(nums)-2, 0, -1):
            if nums[i] >= d:
                d = 1
                continue
            d += 1
        return nums[0] >= d
        '''
        last = len(nums)-1
        for i in xrange(len(nums)-2, -1, -1):
            if i+nums[i] >= last:
                last = i
        return last == 0
        '''
        n = len(nums)
        i = 0
        reach = 0
        while i <= reach and i < n:
            reach = max(i+nums[i], reach)
            i += 1
        return i == n
        '''
