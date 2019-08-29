__author__ = 'Chelsea'
class Solution(object):
    def firstMissingPositive(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        l = len(nums)
        for i in xrange(l):
            while 0 < nums[i] <= l and nums[i] != nums[nums[i]-1]:
                temp = nums[i]
                nums[i], nums[temp-1] = nums[temp-1], nums[i]
        #print nums
        for i in xrange(l):
            if nums[i] != i+1:
                return i+1
        return l+1

test = Solution()
print test.firstMissingPositive([2, 4, -1, 0, 1])