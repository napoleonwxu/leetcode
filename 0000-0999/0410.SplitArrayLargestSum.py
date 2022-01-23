class Solution(object):
    def splitArray(self, nums, m):
        """
        :type nums: List[int]
        :type m: int
        :rtype: int
        """
        low, high = max(nums), sum(nums)
        while low < high:
            mid = low + (high-low)/2
            if self.isVaild(nums, m, mid):
                high = mid
            else:
                low = mid + 1
        return high
        
    def isVaild(self, nums, m, test):
        cur = 0
        count = 1
        for num in nums:
            cur += num
            if cur > test:
                cur = num
                count += 1
                if count > m:
                    return False
        return True
        