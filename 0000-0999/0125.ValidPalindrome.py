class Solution(object):
    def isPalindrome(self, s):
        """
        :type s: str
        :rtype: bool
        """
        if len(s) == 0:
            return True
        s = s.lower()
        l = 0
        r = len(s) - 1
        while l < r:
            if s[l].isalnum() == False:
                l += 1
                continue
            if s[r].isalnum() == False:
                r -= 1
                continue
            if s[l] != s[r]:
                return False
            else:
                l += 1
                r -= 1
        return True
        
        