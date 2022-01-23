class Solution(object):
    def nthUglyNumber(self, n):
        """
        :type n: int
        :rtype: int
        """
        ugly = [1]
        i2, i3, i5 = 0, 0, 0
        for j in range(n-1):
            u2, u3, u5 = 2*ugly[i2], 3*ugly[i3], 5*ugly[i5]
            u = min(u2, u3, u5)
            if u == u2:
                i2 += 1
            if u == u3:
                i3 += 1
            if u == u5:
                i5 += 1
            ugly.append(u)

        return ugly[n-1]
