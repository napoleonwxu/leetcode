import itertools

class Solution(object):
    def maxProduct(self, words):
        """
        :type words: List[str]
        :rtype: int
        """
        ans = 0
        '''
        values = []
        for word in words:
            v = 0
            for ch in set(word):
                v |= 1 << (ord(ch) - 97)
            values.append(v)
        '''
        values = [0] * len(words)
        for i, word in enumerate(words):
            for ch in set(word):
                values[i] |= 1 << (ord(ch) - 97)
        for i, j in itertools.combinations(xrange(len(words)), 2):
            if values[i] & values[j] == 0:
                ans = max(ans, len(words[i])*len(words[j]))
        return ans
