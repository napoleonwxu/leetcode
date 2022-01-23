class Solution(object):
    def canConstruct(self, ransomNote, magazine):
        """
        :type ransomNote: str
        :type magazine: str
        :rtype: bool
        """
        return not collections.Counter(ransomNote) - collections.Counter(magazine)
        dic = {}
        for ch in ransomNote:
            dic[ch] = dic.get(ch, 0) + 1
        for ch in magazine:
            if ch in dic:
                dic[ch] -= 1
                if dic[ch] == 0:
                    dic.pop(ch)
                if not dic:
                    return True
        return not dic
        