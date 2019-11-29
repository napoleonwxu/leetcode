class Solution(object):
    def lengthOfLIS(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        ###Binary Search, O(nlgn)
        lis = [0] * len(nums)
        size = 0
        for num in nums:
            i, j = 0, size
            while i != j:
                m = (i + j) / 2
                if lis[m] < num:
                    i = m + 1
                else:
                    j = m
            lis[i] = num
            if i == size:
                size += 1
        return size

        '''# AC, but what about [1, 10, 20, 2, 3, 4] ?
        lis = 0
        for i in xrange(len(nums)):
            temp = 1
            pre_pre = pre = nums[i]
            for j in xrange(i+1, len(nums)):
                if nums[j] > pre:
                    temp += 1
                    pre_pre = pre
                    pre = nums[j]
                elif pre_pre < nums[j] < pre:
                    pre = nums[j]
            lis = max(lis, temp)
        return lis
        '''
