class Solution(object):
    def restoreIpAddresses(self, s):
        """
        :type s: str
        :rtype: List[str]
        """
        ans = []
        '''
        for i in [1, 2, 3]:
            for j in [i+1, i+2, i+3]:
                for k in [j+1, j+2, j+3]:
                    if k >= len(s):
                        continue
                    s1 = s[:i]
                    s2 = s[i:j]
                    s3 = s[j:k]
                    s4 = s[k:]
                    if self.checkIp([s1, s2, s3, s4]):
                        ans.append(s1 + '.' + s2 + '.' + s3 + '.' + s4)
        '''
        self.dfs(s, 0, 0, '', ans)
        return ans
        
    def dfs(self, s, cnt, index, path, ans):
        if cnt > 4:
            return
        if cnt == 4 and index == len(s):
            ans.append(path)
            return
        for i in xrange(1, 4):
            if index+i > len(s):
                break
            temp = s[index:index+i]
            if temp[0] == '0' and temp != '0' or int(temp) > 255:
                continue
            if cnt < 3:
                temp += '.'
            self.dfs(s, cnt+1, index+i, path+temp, ans)
        
    def checkIp(self, s):
        for si in s:
            if si[0] == '0' and si != '0':
                return False
            if int(si) > 255:
                return False
        return True
        