class Solution(object):
    def uniquePaths(self, m, n):
        """
        :type m: int
        :type n: int
        :rtype: int
        """
        #return math.factorial(m+n-2)/math.factorial(m-1)/math.factorial(n-1)
        '''
        if m <= 0 or n <= 0:
            return
        ans = 1
        small = min(m, n)-1
        big = max(m, n)-1
        for i in xrange(big+small, big, -1):
            ans *= i
        for i in xrange(small, 1, -1):
            ans /= i
        return ans
        '''
        cur = [1]*n
        for i in xrange(1, m):
            for j in xrange(1, n):
                cur[j] += cur[j-1]
        return cur[-1]
