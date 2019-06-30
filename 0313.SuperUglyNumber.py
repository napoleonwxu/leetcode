class Solution(object):
    def nthSuperUglyNumber(self, n, primes):
        """
        :type n: int
        :type primes: List[int]
        :rtype: int
        """
        uglys = [1]
        index = [0] * len(primes)
        temp = primes[:]
        for m in xrange(n-1):
            u = min(temp)
            uglys.append(u)
            for i, t in enumerate(temp):
                if u == t:
                    index[i] += 1
                    temp[i] = uglys[index[i]] * primes[i]
        return uglys[-1]
