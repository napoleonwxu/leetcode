class Solution(object):
    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        if not str:
            return 0
        ans = 0
        s = str.strip()
        sign = 1
        if s[0] == '+':
            s = s[1:]
        elif s[0] == '-':
            sign = -1
            s = s[1:]
        for ch in s:
            n = ord(ch) - ord('0')
            if n >= 0 and n <= 9:
                ans = 10*ans + n
            else:
                break
        ans *= sign
        ans = min(ans, 2**31-1)
        ans = max(ans, -2**31)
        return ans
