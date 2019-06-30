class Solution(object):
    def myPow(self, x, n):
        """
        :type x: float
        :type n: int
        :rtype: float
        """
        '''
        if n == 0:
            return 1
        if n < 0:
            return 1/self.myPow(x, -n)
        if n%2:
            return x*self.myPow(x, n-1)
        return self.myPow(x*x, n/2)
        '''
        if n < 0:
            x = 1/x
            n = -n
        ans = 1
        while n:
            if n&1:
                ans *= x
            x *= x
            n >>= 1
        return ans

        '''
        sign = n < 0
        n = abs(n)
        ans = 1
        while n > 0:
            temp, i = x, 1
            while i <= n:
                ans *= temp
                n -= i
                temp *= x
                i += 1
        if sign:
            return 1/ans
        return ans
        '''