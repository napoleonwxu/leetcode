class Solution:
    # @param {integer[]} nums
    # @return {string}
    def largestNumber(self, nums):
        nums = map(str, nums)
        nums.sort(cmp = lambda x, y: cmp(x+y, y+x), reverse = True)
        # 1st 'cmp': sort() argument
        # 2nd 'cmp': compare function
        return ''.join(nums).lstrip('0') or '0'

test = Solution()
print test.largestNumber([3, 300, 30])