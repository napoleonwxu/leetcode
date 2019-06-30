class Solution(object):
    def largestRectangleArea(self, heights):
        """
        :type heights: List[int]
        :rtype: int
        """
        ans = 0
        heights.append(0)
        stack = [-1]
        for i in xrange(len(heights)):
            while heights[i] < heights[stack[-1]]:
                h = heights[stack.pop()]
                w = i - stack[-1] - 1
                ans = max(ans, h*w)
            stack.append(i)
        return ans

        ''' # TLE
        ans = 0
        length = len(heights)
        for i, val in enumerate(heights):
            left, right = i-1, i+1
            while left >= 0 and heights[left] >= val:
                left -= 1
            while right < length and heights[right] >= val:
                right += 1
            ans = max(ans, val*(right-left-1))
        return ans
        '''