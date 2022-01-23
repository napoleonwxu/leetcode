class Solution(object):
    def findMinHeightTrees(self, n, edges):
        """
        :type n: int
        :type edges: List[List[int]]
        :rtype: List[int]
        """
        if n == 1:
            return [0]
        graph = {i:set() for i in xrange(n)}
        for u, v in edges:
            graph[u].add(v)
            graph[v].add(u)
        '''
        leaves = []
        for i in xrange(n):
            if len(graph[i]) == 1:
                leaves.append(i)
        '''
        leaves = [i for i in xrange(n) if len(graph[i]) == 1]
        while n > 2:
            n -= len(leaves)
            newLeaves = []
            for u in leaves:
                v = graph[u].pop()
                graph[v].remove(u)
                if len(graph[v]) == 1:
                    newLeaves.append(v)
            leaves = newLeaves
        return leaves