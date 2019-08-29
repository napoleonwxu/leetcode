class Solution(object):
    def firstUniqChar(self, s):
        """
        :type s: str
        :rtype: int
        """
        once = {}
        repeat = set()
        for i, ch in enumerate(s):
            if ch not in repeat:
                if ch not in once:
                    once[ch] = i
                else:
                    once.pop(ch)
                    repeat.add(ch)
        return min(once.values()) if once else -1
        