class Solution(object):
    def isIsomorphic(self, s,t):
        ds, dt= {}, {}
        if len(s) != len(t):
            return False
        for i in xrange(len(s)):
            if s[i] in ds:
                ds[s[i]] += [i]
            else:
                ds[s[i]] = [i]
            if t[i] in dt:
                dt[t[i]] += [i]
            else:
                dt[t[i]] = [i]
        return sorted(ds.values()) == sorted(dt.values())

test = Solution()
res = test.isIsomorphic('abba', 'cddc')
print res
