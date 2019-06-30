class Solution(object):
    def calculate(self, s):
        """
        :type s: str
        :rtype: int
        """
        signs = [1, 1]
        i = 0
        ans = 0
        while i < len(s):
            c = s[i]
            if c.isdigit():
                start = i
                while i < len(s) and s[i].isdigit():
                    i += 1
                ans += signs.pop() * int(s[start:i])
                continue
            if c in '(+-':
                #signs += signs[-1] * (1, -1)[c == '-'],
                signs += signs[-1] * [1, -1][c == '-'],
                #signs.append(signs[-1] * [1, -1][c == '-'])
            elif c == ')':
                signs.pop()
            i += 1
        return ans
