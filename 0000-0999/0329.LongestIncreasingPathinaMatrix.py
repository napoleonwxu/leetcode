class Solution(object):
    def longestIncreasingPath(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: int
        """
        def dfs(i, j):
            for x, y in ((i-1, j), (i+1, j), (i, j-1), (i, j+1)):
                if 0 <= x < m and 0 <= y < n and matrix[x][y] > matrix[i][j]:
                    if not step[x][y]:
                        dfs(x, y)
                    step[i][j] = max(step[i][j], step[x][y]+1)
            if not step[i][j]:
                step[i][j] = 1

        if not matrix:
            return 0
        m, n = len(matrix), len(matrix[0])
        step = [[0]*n for i in xrange(m)]
        for i in xrange(len(matrix)):
            for j in xrange(len(matrix[0])):
                if not step[i][j]:
                    dfs(i, j)
        return max([max(t) for t in step])
