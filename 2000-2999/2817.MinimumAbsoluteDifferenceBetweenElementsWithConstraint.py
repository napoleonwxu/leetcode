from sortedcontainers import SortedSet

class Solution:
    def minAbsoluteDifference(self, nums: List[int], x: int) -> int:
        ans = float('inf')
        s = SortedSet()
        for i in range(x, len(nums)):
            if ans == 0:
                break
            s.add(nums[i-x])
            idx = s.bisect(nums[i])
            if idx > 0:
                ans = min(ans, nums[i]-s[idx-1])
            if idx < len(s):
                ans = min(ans, s[idx]-nums[i])
        return ans
