class TrieNode(object):
    def __init__(self):
        self.children = {}
        self.isWord = False

class WordDictionary(object):
    def __init__(self):
        """
        initialize your data structure here.
        """
        self.root = TrieNode()

    def addWord(self, word):
        """
        Adds a word into the data structure.
        :type word: str
        :rtype: void
        """
        node = self.root
        for ch in word:
            if ch not in node.children:
                node.children[ch] = TrieNode()
            node = node.children[ch]
        node.isWord = True

    def search(self, word):
        """
        Returns if the word is in the data structure. A word could
        contain the dot character '.' to represent any one letter.
        :type word: str
        :rtype: bool
        """
        return self.find(self.root, word)

    def find(self, node, word):
        if not word:
            return node.isWord
        if word[0] == '.':
            for ch in node.children:
                if self.find(node.children[ch], word[1:]):
                    return True
        elif word[0] in node.children:
            return self.find(node.children[word[0]], word[1:])
        return False


# Your WordDictionary object will be instantiated and called as such:
wordDictionary = WordDictionary()
wordDictionary.addWord("word")
print wordDictionary.search("pattern")
print wordDictionary.search(".ord")