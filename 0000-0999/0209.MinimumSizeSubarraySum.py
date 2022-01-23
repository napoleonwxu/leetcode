class Solution(object):
    def minSubArrayLen(self, s, nums):
        """
        :type s: int
        :type nums: List[int]
        :rtype: int
        """
        if not s or sum(nums) < s:
            return 0
        minLen = float('inf')
        '''
        i, j = -1, 0
        while i < len(nums):# and j < len(nums):
            if s > 0:
                i += 1
                if i < len(nums):
                    s -= nums[i]
            else:
                minLen = min(minLen, i-j+1)
                s += nums[j]
                j += 1
        '''
        i = j = 0
        sum_t = 0
        while i < len(nums):
            sum_t += nums[i]
            i += 1
            while sum_t >= s:
                minLen = min(minLen, i-j)
                sum_t -= nums[j]
                j += 1
        return minLen
        '''#O(nlogn)
        left, right = 0, len(nums)
        mid = (left+right)/2
        while left <= right:
            if self.judge(mid, s, nums):
                minLen = mid
                right = mid - 1
            else:
                left += 1
        return minLen
        '''

    def judge(self, length, s, nums):
        sums = 0
        for i in xrange(len(nums)):
            sums += nums[i]
            if i >= length:
                sums -= nums[i-length]
            if sums >= s:
                return True
        return False