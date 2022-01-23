class Solution(object):
    def generateMatrix(self, n):
        """
        :type n: int
        :rtype: List[List[int]]
        """
        if n <= 0:
            return []
        if n == 1:
            return [[1]]
        ans = [[0]*n for i in xrange(n)]
        l, r = 0, n-1
        t, b = 0, n-1
        m = 1
        while l < r and t < b:
            for j in xrange(l, r):
                ans[t][j] = m
                m += 1
            for i in xrange(t, b):
                ans[i][r] = m
                m += 1
            for j in xrange(r, l, -1):
                ans[b][j] = m
                m += 1
            for i in xrange(b, t, -1):
                ans[i][l] = m
                m += 1
            l += 1
            r -= 1
            t += 1
            b -= 1
        if l == r:
            ans[l][l] = m
        return ans
