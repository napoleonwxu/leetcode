class Solution(object):
    def fractionToDecimal(self, numerator, denominator):
        """
        :type numerator: int
        :type denominator: int
        :rtype: str
        """
        if numerator == 0:
            return '0'
        ans = ''
        if (numerator < 0) ^ (denominator < 0):
            ans += '-'
        n, d = abs(numerator), abs(denominator)
        ans += str(n/d)
        if n%d == 0:
            return ans
        ans += '.'
        dic = {}
        r = n%d
        while r:
            if r in dic:
                ans = ans[:dic[r]] + '(' + ans[dic[r]:] + ')'
                break
            dic[r] = len(ans)
            r *= 10
            ans += str(r/d)
            r %= d
        return ans