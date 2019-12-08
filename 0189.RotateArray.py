class Solution(object):
    def rotate(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        le = len(nums)
        k %= le
        if k == 0:
            return        
        #nums[:] = nums[-k:] + nums[:-k]
        nums[:] = self.reverse(nums, 0, le-k-1)
        nums[:] = self.reverse(nums, le-k, le-1)
        nums[:] = self.reverse(nums, 0, le-1)
        '''
        temp = []
        for i in xrange(k):
            temp.append(nums[le-k+i])
        for i in xrange(le-k):
            nums[le-i-1] = nums[le-k-i-1]
        for i in xrange(k):
            nums[i] = temp[i]
        '''
    def reverse(self, nums, i , j):
        while i < j:
            nums[i], nums[j] = nums[j], nums[i]
            i += 1
            j -= 1
        return nums