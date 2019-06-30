class Solution(object):
    def jump(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        reach = 0
        step = 0
        temp = 0
        for i in xrange(len(nums)-1):
            reach = max(i+nums[i], reach)
            if i == temp:
                step += 1
                temp = reach
        return step


