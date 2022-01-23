class Solution(object):
    def addStrings(self, num1, num2):
        """
        :type num1: str
        :type num2: str
        :rtype: str
        """
        ans = ''
        i, j = len(num1)-1, len(num2)-1
        c = 0
        while i >= 0 or j >= 0 or c:
            if i >= 0:
                c += int(num1[i])
                i -= 1
            if j >= 0:
                c += int(num2[j])
                j -= 1
            ans = str(c%10) + ans
            c /= 10
        return ans
        