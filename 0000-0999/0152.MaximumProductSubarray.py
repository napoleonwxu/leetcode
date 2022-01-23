class Solution(object):
    def maxProduct(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if not nums:
            return
        ans = big = small = nums[0]
        for num in nums[1:]:
            big, small = max(num, num*big, num*small), min(num, num*big, num*small)
            ans = max(ans, big)
        return ans

print Solution().maxProduct([-4, -3, -2])