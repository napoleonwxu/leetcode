class Solution(object):
    def climbStairs(self, n):
        """
        :type n: int
        :rtype: int
        """
        '''
        if n < 1:
            return 0
        sum = 0
        for k in xrange(n/2+1):
            m = n-k
            t = 1
            for c in xrange(k):
                t *= m-c
            for c in xrange(2, k+1):
                t /= c
            sum += t
        return sum
        '''
        if n < 3:
            return n
        fir, sec = 1, 2
        for i in xrange(n-2):
            sec, fir = sec+fir, sec
        return sec
        