class Solution(object):
    def grayCode(self, n):
        """
        :type n: int
        :rtype: List[int]
        """
        return [i^(i>>1) for i in xrange(2**n)]
        if n == 0:
            return [0]
        '''
        if n == 1:
            return [0, 1]
        '''
        half1 = self.grayCode(n-1)
        add = 2**(n-1)
        half2 = [m+add for m in reversed(half1)]
        return half1+half2
        