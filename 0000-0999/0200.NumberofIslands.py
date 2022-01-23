class Solution(object):
    def numIslands(self, grid):
        """
        :type grid: List[List[str]]
        :rtype: int
        """
        if not grid:
            return 0
        row = len(grid)
        col = len(grid[0])
        def judge(i, j):
            if 0 <= i < row and 0 <= j < col and grid[i][j] == '1':
                grid[i][j] = '0'
                map(judge, (i, i, i-1, i+1), (j-1, j+1, j, j))
                return 1
            return 0
        return sum(judge(i, j) for i in xrange(row) for j in xrange(col))
        '''
        self.grid = grid
        ans = 0
        for i in xrange(len(self.grid)):
            for j in xrange(len(self.grid[i])):
               ans += self.judge(i, j)
        return ans
        '''

    def judge(self, i, j):
        if 0 <= i < len(self.grid) and 0 <= j < len(self.grid[i]) and self.grid[i][j] == '1':
            self.grid[i][j] = '0'
            map(self.judge, (i, i, i-1, i+1), (j-1, j+1, j, j))
            return 1
        return 0
