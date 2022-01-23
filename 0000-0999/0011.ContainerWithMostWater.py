class Solution(object):
    def maxArea(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        ans = 0
        left, right = 0, len(height)-1
        '''
        while left < right:
            ans = max(ans, min(height[left], height[right])*(right-left))
            if height[left] <= height[right]:
                left += 1
            else:
                right -= 1
        '''
        while left < right:
            h = min(height[left], height[right])
            ans = max(ans, h*(right-left))
            while left < right and height[left] <= h:
                left += 1
            while left < right and height[right] <= h:
                right -= 1
        return ans
        