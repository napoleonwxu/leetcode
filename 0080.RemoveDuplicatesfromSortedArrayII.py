class Solution(object):
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        '''
        length = len(nums)
        if length <= 2:
            return length
        pre = 2
        for num in nums[2:]:
            if num != nums[pre-2]:
                nums[pre] = num
                pre += 1
        return pre
        '''
        i = 0
        for num in nums:
            if i < 2 or num != nums[i-2]:
                nums[i] = num
                i += 1
        return i