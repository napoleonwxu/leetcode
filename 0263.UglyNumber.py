class Solution(object):
    def isUgly(self, num):
        """
        :type num: int
        :rtype: bool
        """
        if num < 1:
            return False
        for m in (2, 3, 5):
            while num % m == 0:
                num /= m
        return num == 1
