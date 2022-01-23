class Solution(object):
    def maxSlidingWindow(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: List[int]
        """
        if len(nums) < k or k == 0:
            return []
        queue = []
        ans = []
        for i in xrange(len(nums)):
            while queue and nums[queue[-1]] < nums[i]:
                queue.pop()
            queue.append(i)
            if i < k-1:
                continue
            if queue and queue[0] <= i-k:
                queue.pop(0)
            ans.append(nums[queue[0]])
        return ans
