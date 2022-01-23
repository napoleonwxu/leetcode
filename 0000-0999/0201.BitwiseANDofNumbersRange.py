class Solution(object):
    def rangeBitwiseAnd(self, m, n):
        """
        :type m: int
        :type n: int
        :rtype: int
        """
        count = 0
        while m != n:
            count += 1
            m >>= 1
            n >>= 1
        return m*(2**count)
        #return m << count
        '''
        fac = 1
        while m != n:
            m >>= 1
            n >>= 1
            fac <<= 1
        return m*fac
        '''