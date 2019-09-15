class Solution(object):
    def mySqrt(self, x):
        """
        :type x: int
        :rtype: int
        """
        #2: Integer Newton
        r = x
        while r*r > x:
            r = (r + x/r)/2
        return r
        
        '''#1: Binary Search
        l, r = 1, x
        while l <= r:
            m = l + (r-l)/2
            if m*m == x:
                return m
            elif m*m < x:
                l = m + 1
            else:
                r = m - 1
        return r
        '''