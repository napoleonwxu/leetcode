class Solution(object):
    def moveZeroes(self, nums):
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        i = 0
        '''
        for j in xrange(len(nums)):
            if nums[j] != 0:
                nums[i] = nums[j]
                i += 1
        for i in xrange(i, len(nums)):
            nums[i] = 0
        '''
        for j in xrange(len(nums)):
            if nums[j] != 0:
                nums[i], nums[j] = nums[j], nums[i]
                i += 1