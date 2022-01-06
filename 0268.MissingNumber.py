class Solution(object):
    def missingNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        return (1+len(nums))*len(nums)/2 - sum(nums)
        return sum(range(len(nums)+1)) - sum(nums)
        return reduce(operator.xor, nums + [i for i in xrange(len(nums)+1)])
        return reduce(operator.xor, nums + range(len(nums)+1))
        '''
        nums.sort()
        for i in xrange(len(nums)):
            if i != nums[i]:
                return i
        return len(nums)
        '''