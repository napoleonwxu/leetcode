class Solution(object):
    def findPeakElement(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        '''
        for i in xrange(len(nums)-1):
            if nums[i] > nums[i+1]:
                return i
        return len(nums)-1
        '''
        '''
        nums = [float('-inf')] + nums + [float('-inf')]
        for i in xrange(1, len(nums)-1):
            if nums[i-1] < nums[i] > nums[i+1]:
                return i-1
        '''
        # O(lgn)
        left, right = 0, len(nums)-1
        while left < right:
            mid = (left+right)/2
            if nums[mid] > nums[mid+1]:
                right = mid
            else:
                left = mid + 1
        return left