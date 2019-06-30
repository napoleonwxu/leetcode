# Definition for a undirected graph node
# class UndirectedGraphNode(object):
#     def __init__(self, x):
#         self.label = x
#         self.neighbors = []

class Solution(object):
    def cloneGraph(self, node):
        """
        :type node: UndirectedGraphNode
        :rtype: UndirectedGraphNode
        """
        if not node:
            return None
        self.dic = {}
        return self.createNode(node)

    def createNode(self, node):
        new = UndirectedGraphNode(node.label)
        self.dic[new.label] = new
        for n in node.neighbors:
            if n.label not in self.dic:
                self.createNode(n)
            new.neighbors.append(self.dic[n.label])
        return new