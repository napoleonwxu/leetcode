class Solution(object):
    def isScramble(self, s1, s2):
        """
        :type s1: str
        :type s2: str
        :rtype: bool
        """
        m, n = len(s1), len(s2)
        if m != n or sorted(s1) != sorted(s2):
            return False
        if m <= 3 or s1 == s2:
            return True
        f = self.isScramble
        for i in xrange(1, n):
            if f(s1[:i], s2[:i]) and f(s1[i:], s2[i:]) or f(s1[:i], s2[-i:]) and f(s1[i:], s2[:-i]):
                return True
        '''
        for i in xrange(1, n):
            if self.isScramble(s1[:i], s2[:i]) and self.isScramble(s1[i:], s2[i:]) or \
               self.isScramble(s1[:i], s2[-i:]) and self.isScramble(s1[i:], s2[:-i]):
                return True
        '''
        return False