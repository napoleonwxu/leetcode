class Solution(object):
    def productExceptSelf(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        length = len(nums)
        left, right = [1]*length, [1]*length
        for i in xrange(length-1):
            left[i+1] = nums[i] * left[i]
        for i in xrange(length-1, 0, -1):
            right[i-1] = nums[i] * right[i]
        return map(lambda x,y:x*y, left, right)