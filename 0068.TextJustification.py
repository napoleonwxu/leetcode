class Solution(object):
    def fullJustify(self, words, maxWidth):
        """
        :type words: List[str]
        :type maxWidth: int
        :rtype: List[str]
        """
        ans = []
        cur, temp = [], 0
        for word in words:
            if temp + len(cur) + len(word) > maxWidth:
                for i in xrange(maxWidth-temp):
                    cur[i%(len(cur)-1 or 1)] += ' '
                ans.append(''.join(cur))
                cur, temp = [], 0
            cur.append(word)
            temp += len(word)
        return ans+[' '.join(cur).ljust(maxWidth)]
