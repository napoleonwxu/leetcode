class Solution(object):
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        return 2*sum(set(nums)) - sum(nums)
        return reduce(operator.xor, nums)
        ans = 0
        for num in nums:
            ans ^= num
        return ans
        