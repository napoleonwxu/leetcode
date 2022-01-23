class Solution(object):
    def partition(self, s):
        """
        :type s: str
        :rtype: List[List[str]]
        """
        ans = []
        self.dfs(s, [], ans)
        return ans

    def dfs(self, s, path, ans):
        if not s:
            ans.append(path)
            return
        for i in xrange(1, len(s)+1):
            if self.isPal(s[:i]):
                self.dfs(s[i:], path+[s[:i]], ans)

    def isPal(self, s):
        return s == s[::-1]