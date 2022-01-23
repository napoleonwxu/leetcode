class Solution(object):
    def multiply(self, num1, num2):
        """
        :type num1: str
        :type num2: str
        :rtype: str
        """
        mul = [0] * (len(num1)+len(num2))
        pos = len(mul) - 1
        for d1 in reversed(num1):
            index = pos
            for d2 in reversed(num2):
                mul[index] += int(d1) * int(d2)
                mul[index-1] += mul[index]/10
                mul[index] %= 10
                index -= 1
            pos -= 1
        pos = 0
        while pos < len(mul)-1 and mul[pos] == 0:
            pos += 1
        ans = ''
        for i in xrange(pos, len(mul)):
            ans += str(mul[i])
        return ans
