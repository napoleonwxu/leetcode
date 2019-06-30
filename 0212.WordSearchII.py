class TrieNode(object):
    def __init__(self):
        self.children = {}
        self.isWord = False

class Trie(object):
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                node.children[ch] = TrieNode()
            node = node.children[ch]
        node.isWord = True

    def search(self, word):
        node = self.root
        for ch in word:
            if ch not in node.children:
                return False
            node = node.children[ch]
        return node.isWord

class Solution(object):
    def findWords(self, board, words):
        """
        :type board: List[List[str]]
        :type words: List[str]
        :rtype: List[str]
        """
        ans = []
        trie = Trie()
        for word in words:
            trie.insert(word)
        for i in xrange(len(board)):
            for j in xrange(len(board[0])):
                self.dfs(board, trie.root, i, j, '', ans)
        return ans

    def dfs(self, board, node, i, j, path, ans):
        if node.isWord:
            ans.append(path)
            node.isWord = False
        if i < 0 or i >= len(board) or j < 0 or j >= len(board[0]):
            return
        temp = board[i][j]
        if temp not in node.children:
            return
        board[i][j] = '#'
        node = node.children[temp]
        self.dfs(board, node, i-1, j, path+temp, ans)
        self.dfs(board, node, i+1, j, path+temp, ans)
        self.dfs(board, node, i, j-1, path+temp, ans)
        self.dfs(board, node, i, j+1, path+temp, ans)
        board[i][j] = temp
