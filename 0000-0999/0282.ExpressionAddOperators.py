class Solution(object):
    def addOperators(self, num, target):
        """
        :type num: str
        :type target: int
        :rtype: List[str]
        """
        if not num:
            return []
        ans = []
        self.dfs(num, target, '', 0, 0, 0, ans)
        return ans

    def dfs(self, num, target, path, pos, val, multed, ans):
        if pos == len(num) and val == target:
            ans.append(path)
            return
        for i in xrange(pos, len(num)):
            if num[pos] == '0' and i != pos:
                break
            t = num[pos:i+1]
            cur = int(t)
            if pos == 0:
                self.dfs(num, target, t, i+1, cur, cur, ans)
            else:
                self.dfs(num, target, path + '+' + t, i+1, val+cur, cur, ans)
                self.dfs(num, target, path + '-' + t, i+1, val-cur, -cur, ans)
                self.dfs(num, target, path + '*' + t, i+1, val-multed+cur*multed, cur*multed, ans)
