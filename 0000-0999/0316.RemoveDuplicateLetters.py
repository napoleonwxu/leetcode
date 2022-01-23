class Solution(object):
    def removeDuplicateLetters(self, s):
        """
        :type s: str
        :rtype: str
        """
        ans = ''
        for i, ch in enumerate(s):
            if ch not in ans:
                while ch < ans[-1:] and i < s.rindex(ans[-1]):
                    ans = ans[:-1]
                ans += ch
        return ans
        '''
        # recursively
        sigle = set(s)
        for ch in sorted(sigle):
            sub = s[s.index(ch):]
            if set(sub) == sigle:
                return ch + self.removeDuplicateLetters(sub.replace(ch, ''))
        return ''
        '''