class Solution(object):
    def maximalRectangle(self, matrix):
        """
        :type matrix: List[List[str]]
        :rtype: int
        """
        if not matrix:
            return 0
        m = len(matrix)
        n = len(matrix[0])
        maxRe = 0
        left, right = [0] * n, [n] * n
        height = [0] * n
        for i in xrange(m):
            cur_left, cur_right = 0, n
            for j in xrange(n):
                if matrix[i][j] == '1':
                    height[j] += 1
                    left[j] = max(left[j], cur_left)
                else:
                    height[j] = 0
                    left[j] = 0
                    cur_left = j+1
            for j in xrange(n-1, -1, -1):
                if matrix[i][j] == '1':
                    right[j] = min(right[j], cur_right)
                else:
                    right[j] = n
                    cur_right = j
            for j in xrange(n):
                maxRe = max(maxRe, (right[j]-left[j])*height[j])
        return maxRe
