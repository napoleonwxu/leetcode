class Solution(object):
    def minWindow(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: str
        """
        if not s or not t or len(s) < len(t):
            return ""
        require = {}
        for ch in t:
            require[ch] = require.get(ch, 0) + 1
        #chSet = set(t)
        i, j = -1, 0
        count = len(t)
        minLen = float('inf')
        while i < len(s) and j < len(s):
            if count:
                i += 1
                if i < len(s):
                    require[s[i]] = require.get(s[i], 0) - 1
                    if s[i] in t and require.get(s[i]) >= 0:
                        count -= 1
            else:
                if minLen > i-j+1:
                    minLen = i-j+1
                    minIdx = j
                require[s[j]] = require.get(s[j], 0) + 1
                if s[i] in t and require.get(s[j]) > 0:
                    count += 1
                j += 1
        if minLen == float('inf'):
            return ""
        else:
            return s[minIdx:minIdx+minLen]
