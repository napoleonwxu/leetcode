class Solution(object):
    def findRepeatedDnaSequences(self, s):
        """
        :type s: str
        :rtype: List[str]
        """
        words = set()
        repeat = set()
        for i in xrange(len(s)-9):
            temp = s[i:i+10]
            if temp in words:
                repeat.add(temp)
            else:
                words.add(temp)
        return list(repeat)
