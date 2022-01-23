class Solution(object):
    def findDuplicate(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if len(nums) <= 1:
            return
        slow = nums[0]
        fast = nums[nums[0]]
        while slow != fast:
            slow = nums[slow]
            fast = nums[nums[fast]]
        slow = 0
        while slow != fast:
            slow = nums[slow]
            fast = nums[fast]
        return slow
        
        '''
        dic = {}
        for num in nums:
            dic[num] = dic.get(num, 0) + 1
        for num in dic.keys():
            if dic[num] != 1:
                return num
        '''    