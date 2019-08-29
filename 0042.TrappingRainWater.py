class Solution(object):
    def trap(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        ans = 0
        left, right = 0, len(height)-1
        maxLeft = maxRight = 0
        while left < right:
            if height[left] <= height[right]:
                maxLeft = max(maxLeft, height[left])
                ans += maxLeft - height[left]
                left += 1
            else:
                maxRight = max(maxRight, height[right])
                ans += maxRight - height[right]
                right -= 1
        return ans
                    
        