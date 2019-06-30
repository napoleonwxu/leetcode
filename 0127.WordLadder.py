class Solution(object):
    def ladderLength(self, beginWord, endWord, wordList):
        """
        :type beginWord: str
        :type endWord: str
        :type wordList: Set[str]
        :rtype: int
        """
        queue = [(beginWord, 1)]
        visited = set()
        wordLen = len(beginWord)
        while queue:
            word, step = queue.pop(0)
            if word == endWord:
                return step
            for i in xrange(wordLen):
                for ch in 'abcdefghijklmnopqrstuvwxyz':
                    new = word[:i] + ch + word[i+1:]
                    if new in wordList and new not in visited:
                    #if (new in wordList or new == endWord) and new not in visited:
                        queue.append((new, step+1))
                        visited.add(new)
        return 0
