class Solution(object):
    def longestSubstring(self, s, k):
        """
        :type s: str
        :type k: int
        :rtype: int
        """
        for ch in set(s):
            if s.count(ch) < k:
                return max(self.longestSubstring(sub, k) for sub in s.split(ch))
        return len(s)
        '''# recursive
        if len(s) < k:
            return 0
        dic = {}
        for ch in s:
            dic[ch] = dic.get(ch, 0) + 1
        index = 0
        while index < len(s) and dic[s[index]] >= k:
            index += 1
        if index == len(s):
            return index
        l = self.longestSubstring(s[:index], k)
        r = self.longestSubstring(s[index+1:], k)
        return max(l, r)
        '''