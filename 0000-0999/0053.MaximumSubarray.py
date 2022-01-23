class Solution(object):
    def maxSubArray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        #return self.divide(nums, 0, len(nums)-1)   # O(nlogn)
        # O(n)
        pre = ans = nums[0]
        for num in nums[1:]:
            pre = max(num, pre+num)
            ans = max(pre, ans)
        return ans
    
    def divide(self, nums, l, r):
        if l >= r:
            return nums[l]
        m = l + (r-l)/2
        maxl = suml = nums[m]
        for i in xrange(m-1, l-1, -1):
            suml += nums[i]
            maxl = max(maxl, suml)
        maxr = sumr = nums[m+1]
        for i in xrange(m+2, r+1):
            sumr += nums[i]
            maxr = max(maxr, sumr)
        return max(maxl+maxr, self.divide(nums, l, m), self.divide(nums, m+1, r))
        