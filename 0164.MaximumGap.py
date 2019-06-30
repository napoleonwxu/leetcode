import math

class Solution(object):
    def maximumGap(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        length = len(nums)
        if length < 2:
            return 0
        minNum, maxNum = min(nums), max(nums)
        gap = int(math.ceil(float(maxNum-minNum)/(length-1)))
        bucketsMin = [float('inf')] * (length-1)
        bucketsMax = [float('-inf')] * (length-1)
        for num in nums:
            if num == minNum or num == maxNum:
                continue
            i = (num-minNum)/gap
            bucketsMin[i] = min(bucketsMin[i], num)
            bucketsMax[i] = max(bucketsMax[i], num)
        print bucketsMin, bucketsMax
        ans = float('-inf')
        pre = minNum
        for i in xrange(length-1):
            if bucketsMin[i] == float('inf') and bucketsMax[i] == float('-inf'):
                continue
            ans = max(ans, bucketsMin[i]-pre)
            pre = bucketsMax[i]
        ans = max(ans, maxNum-pre)
        return ans

        ''' # O(NlgN)
        nums.sort()
        ans = 0
        for i in xrange(1, len(nums)):
            ans = max(ans, nums[i]-nums[i-1])
        return ans
        '''

print Solution().maximumGap([3, 6, 9, 1])