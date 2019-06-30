import collections
import string

class Solution(object):
    def findLadders(self, beginWord, endWord, wordlist):
        """
        :type beginWord: str
        :type endWord: str
        :type wordlist: Set[str]
        :rtype: List[List[int]]
        """
        wordlist.add(endWord)
        wordLen = len(beginWord)
        level = {beginWord}
        parents = collections.defaultdict(set)
        while level and endWord not in parents:
            nextLevel = collections.defaultdict(set)
            for word in level:
                for ch in string.ascii_lowercase:
                    for i in xrange(wordLen):
                        temp = word[:i] + ch + word[i+1:]
                        if temp in wordlist and temp not in parents:
                            nextLevel[temp].add(word)
            level = nextLevel
            parents.update(nextLevel)
        res = [[endWord]]
        while res and res[0][0] != beginWord:
            res = [[p]+r for r in res for p in parents[r[0]]]
        return res
        '''
        dic.add(end)
        level = {start}
        parents = collections.defaultdict(set)
        while level and end not in parents:
            next_level = collections.defaultdict(set)
            for node in level:
                for char in string.ascii_lowercase:
                    for i in range(len(start)):
                        n = node[:i]+char+node[i+1:]
                        if n in dic and n not in parents:
                            next_level[n].add(node)
            level = next_level
            parents.update(next_level)
        res = [[end]]
        while res and res[0][0] != start:
            res = [[p]+r for r in res for p in parents[r[0]]]
        return res
        '''
