class Solution(object):
    def maximalSquare(self, matrix):
        """
        :type matrix: List[List[str]]
        :rtype: int
        """
        if not matrix:
            return 0
        row = len(matrix) + 1
        col = len(matrix[0]) + 1
        dp = [[0]*col for i in xrange(row)]
        maxLen = 0
        for i in xrange(1, row):
            for j in xrange(1, col):
                if matrix[i-1][j-1] == '1':
                    dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
                    maxLen = max(maxLen, dp[i][j])
        return maxLen**2
