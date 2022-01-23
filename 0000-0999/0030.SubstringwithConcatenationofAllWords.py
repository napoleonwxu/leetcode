class Solution(object):
    def findSubstring(self, s, words):
        """
        :type s: str
        :type words: List[str]
        :rtype: List[int]
        """
        if not s or not words or not words[0]:
            return []
        n = len(s)
        k = len(words[0])
        t = len(words)*k
        ans = []
        req = {}
        for word in words:
            req[word] = req.get(word, 0) + 1
        for i in xrange(min(k, n-t+1)):
            self.subString(i, i, n, k, t, s, req, ans)
        return ans

    def subString(self, l, r, n, k, t, s, req, ans):
        temp = {}
        while r+k <= n:
            word = s[r:r+k]
            r += k
            if word not in req:
                l = r
                temp.clear()
            else:
                temp[word] = temp.get(word, 0) + 1
                while temp[word] > req[word]:
                    temp[s[l:l+k]] -= 1
                    l += k
                if r-l == t:
                    ans.append(l)
