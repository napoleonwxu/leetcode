class Solution(object):
    def uniquePathsWithObstacles(self, obstacleGrid):
        """
        :type obstacleGrid: List[List[int]]
        :rtype: int
        """
        m = len(obstacleGrid)
        n = len(obstacleGrid[0])
        '''
        for j in xrange(n):
            if obstacleGrid[0][j] == 1:
                obstacleGrid[0][j:] = [0]*(n-j)
                break
            obstacleGrid[0][j] = 1
        for i in xrange(1, m):
            if obstacleGrid[i][0] == 1:
                obstacleGrid[i][0] = 0
            else:
                obstacleGrid[i][0] = obstacleGrid[i-1][0]
            for j in xrange(1, n):
                if obstacleGrid[i][j] == 1:
                    obstacleGrid[i][j] = 0
                else:
                    obstacleGrid[i][j] = obstacleGrid[i][j-1] + obstacleGrid[i-1][j]
        return obstacleGrid[-1][-1]
        '''
        dp = [0]*n
        dp[0] = 1
        for row in obstacleGrid:
            for j in xrange(n):
                if row[j] == 1:
                    dp[j] = 0
                elif j > 0:
                    dp[j] += dp[j-1]
        return dp[-1]
