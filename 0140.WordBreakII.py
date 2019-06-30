class Solution(object):
    def wordBreak(self, s, wordDict):
        """
        :type s: str
        :type wordDict: Set[str]
        :rtype: List[str]
        """
        return self.findWords(0, len(s), s, wordDict, {})
        '''
        ans = []
        self.dfs(0, len(s), s, wordDict, '', ans)
        return ans
        '''

    def findWords(self, start, end, s, wordDict, catch):
        if start in catch:
            return catch[start]
        catch[start] = []
        candidate = ''
        cur = start
        while cur < end:
            candidate += s[cur]
            cur += 1
            if candidate in wordDict:
                if cur == end:
                    catch[start].append(candidate)
                else:
                    for x in self.findWords(cur, end, s, wordDict, catch):
                        catch[start].append(candidate + ' ' + x)
        return catch[start]

    def dfs(self, start, end, s, wordDict, temp, ans):  #TLE
        if start >= end:
            ans.append(temp)
        for i in xrange(start, end):
            if s[start:i+1] in wordDict:
                if start:
                    self.dfs(i+1, end, s, wordDict, temp + ' ' + s[start:i+1], ans)
                else:
                    self.dfs(i+1, end, s, wordDict, s[start:i+1], ans)
