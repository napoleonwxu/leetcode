class Solution(object):
    def isNumber(self, s):
        """
        :type s: str
        :rtype: bool
        """
        s = s.strip()
        if not s:
            return False
        sign = digit = point = e = False
        for ch in s:
            if ch >= '0' and ch <= '9':
                sign = digit = True
            elif ch == '.' and not point:
                sign = point = True
            elif ch in '+-' and not sign and not digit:
                sign = True
            elif ch in 'eE' and not e and digit:
                e = point = True
                sign = digit = False
            else:
                return False
        if digit:
            return True
        else:
            return False
        