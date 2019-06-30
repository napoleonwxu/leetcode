import itertools

class Solution(object):
    def isAdditiveNumber(self, num):
        """
        :type num: str
        :rtype: bool
        """
        n = len(num)
        for i, j in itertools.combinations(xrange(1, n), 2):
            if i > 1 and num[0] == '0': #0235813 is False
                break
            a, b = num[:i], num[i:j]
            if b != str(int(b)):
                continue
            while j < n:
                c = str(int(a) + int(b))
                if not num.startswith(c, j):
                    break
                a, b = b, c
                j += len(c)
            if j == n:
                return True
        return False

