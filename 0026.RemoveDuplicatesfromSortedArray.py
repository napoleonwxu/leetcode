class Solution(object):
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if not nums:
            return 0
        i = 0
        for num in nums[1:]:
            if nums[i] != num:
                i += 1
                nums[i] = num
        return i+1
        '''
        l = len(nums)
        count = 0
        for i in xrange(1, l):
            if nums[i] == nums[i-1]:
                count += 1
            else:
                nums[i-count] = nums[i]
        return l-count
        '''
