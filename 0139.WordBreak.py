class Solution(object):
    def wordBreak(self, s, wordDict):
        """
        :type s: str
        :type wordDict: Set[str]
        :rtype: bool
        """
        isTrue = [True] + [False]*len(s)
        for i in xrange(len(s)):
            for j in xrange(i, len(s)):
                if isTrue[i] and s[i:j+1] in wordDict:
                    isTrue[j+1] = True
        return isTrue[-1]
        '''
        isTrue = [False] * (len(s)+1)
        for i in xrange(len(s)):
            for j in xrange(i+1, len(s)+1):
                if s[i:j] in wordDict and (isTrue[i] or i == 0):
                    isTrue[j] = True
        return isTrue[-1]
        '''
        '''
        isTrue = [False] * len(s)
        for i in xrange(len(s)):
            for word in wordDict:
                if word == s[i-len(word)+1:i+1] and (i-len(word) < 0 or isTrue[i-len(word)]):
                    isTrue[i] = True
        return isTrue[-1]
        '''
