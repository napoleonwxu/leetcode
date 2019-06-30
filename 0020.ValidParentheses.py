class Solution(object):
    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        q = {')':'(', ']':'[', '}':'{'}
        d = []
        for ch in s:
            if ch in q.keys():
                if not d or d.pop() != q[ch]:
                    return False
            else:
                d.append(ch)
        if d:
            return False
        return True
