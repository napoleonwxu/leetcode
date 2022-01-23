class Solution(object):
    def calculate(self, s):
        """
        :type s: str
        :rtype: int
        """
        stack = []
        num = 0
        preSign = '+'
        for i in xrange(len(s)):
            if s[i].isdigit():
                num = 10*num + int(s[i])
            if s[i] in '+-*/' or i == len(s)-1:
                if preSign == '+':
                    stack.append(num)
                elif preSign == '-':
                    stack.append(-num)
                elif preSign == '*':
                    stack.append(stack.pop()*num)
                elif preSign == '/' and num != 0:
                    stack.append(int(stack.pop()/float(num)))
                num = 0
                preSign = s[i]
        return sum(stack)
