class Solution(object):
    def findAnagrams(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: List[int]
        """
        ans = []
        count = [0] * 128
        for ch in p:
            count[ord(ch)] += 1
        left = right = 0
        c = n = len(p)
        while right < len(s):
            if count[ord(s[right])] > 0:
                c -= 1
            count[ord(s[right])] -= 1
            right += 1
            
            if c == 0:
                ans.append(left)
            
            if right - left == n:
                if count[ord(s[left])] >= 0:
                    c += 1
                count[ord(s[left])] += 1
                left += 1
        return ans
        