class Solution(object):
    def totalNQueens(self, n):
        """
        :type n: int
        :rtype: int
        """
        self.ans = 0
        self.dfs(n, [], [], [])
        return self.ans

    def dfs(self, n, cols, xySum, xyDif):
        p = len(cols)
        if p == n:
            self.ans += 1
            return
        for q in xrange(n):
            if q not in cols and p+q not in xySum and p-q not in xyDif:
                self.dfs(n, cols+[q], xySum+[p+q], xyDif+[p-q])


test = Solution()
for i in xrange(1, 11):
    print test.totalNQueens(i)