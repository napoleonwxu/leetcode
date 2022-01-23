class Solution(object):
    def trailingZeroes(self, n):
        """
        :type n: int
        :rtype: int
        """
        #return 0 if n == 0 else n/5 + self.trailingZeroes(n/5)
        #return n and n/5 + self.trailingZeroes(n/5)
        ans = 0
        ''' #TLE
        for i in xrange(5, n+1, 5):
            tmp = i
            while tmp%5 == 0:
                ans += 1
                tmp /= 5
        '''
        while n:
            n /= 5
            ans += n
        return ans
