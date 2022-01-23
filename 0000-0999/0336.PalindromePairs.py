class Solution(object):
    def palindromePairs(self, words):
        """
        :type words: List[str]
        :rtype: List[List[int]]
        """
        ans = []
        dic = {word: i for i, word in enumerate(words)}
        for i, word in enumerate(words):
            for j in xrange(len(word)+1):
                p1, p2 = word[:j], word[j:]
                if p1[::-1] in dic and dic[p1[::-1]] != i and p2 == p2[::-1]:
                    ans.append([i, dic[p1[::-1]]])
                if j != 0 and p2[::-1] in dic and dic[p2[::-1]] != i and p1 == p1[::-1]:
                    ans.append([dic[p2[::-1]], i])
        return ans