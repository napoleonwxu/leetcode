class Solution(object):
    def titleToNumber(self, s):
        """
        :type s: str
        :rtype: int
        """
        res = 0
        for char in s:
            res = res*26 + ord(char) - ord('A') + 1
        return res
        
        