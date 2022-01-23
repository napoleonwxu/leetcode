class Solution(object):
    def solveNQueens(self, n):
        """
        :type n: int
        :rtype: List[List[str]]
        """
        ans = []
        self.dfs(n, [], [], [], ans)
        return [['.'*col + 'Q' + '.'*(n-col-1) for col in cols] for cols in ans]

    def dfs(self, n, cols, xySum, xyDif, ans):
        p = len(cols)
        if p == n:
            ans.append(cols)
            return
        for q in xrange(n):
            if q not in cols and p+q not in xySum and p-q not in xyDif:
                self.dfs(n, cols+[q], xySum+[p+q], xyDif+[p-q], ans)

test = Solution()
print test.solveNQueens(4)
