import math

class Solution(object):
    def countPrimes(self, n):
        """
        :type n: int
        :rtype: int
        """
        if n < 2:
            return 0
        isPrime = [True] * n
        isPrime[0] = isPrime[1] = False
        for i in xrange(4, n, 2):
            isPrime[i] = False
        for i in xrange(3, int(math.sqrt(n))+1, 2):
            if isPrime[i]:
                i2 = i + i
                j = i * i
                while j < n:
                    isPrime[j] = False
                    j += i2
        return isPrime.count(True)
