class Solution(object):
    def longestConsecutive(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        dic = {}
        ans = 0
        for num in nums:
            if num not in dic:
                left = dic.get(num-1, 0)
                right = dic.get(num+1, 0)
                new = left + right + 1
                ans = max(ans, new)
                dic[num] = dic[num-left] = dic[num+right] = new
            print dic
        return ans

test = Solution()
print test.longestConsecutive([100, 4, 200, 1, 3, 2])
