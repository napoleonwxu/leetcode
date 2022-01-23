class Solution(object):
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        #return (-1 if x < 0 else 1) * int(str(abs(x))[::-1])
        ans = 0
        sign = 1
        if x < 0:
            sign = -1
            x *= -1
        while x:
            ans = 10*ans + x%10
            x /= 10
        ans *= sign
        if ans > 2**31-1 or ans < -2**31:
            return 0
        return ans
