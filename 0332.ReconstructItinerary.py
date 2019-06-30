from collections import defaultdict

class Solution(object):
    def findItinerary(self, tickets):
        """
        :type tickets: List[List[str]]
        :rtype: List[str]
        """
        targets = {}
        for f, t in sorted(tickets)[::-1]:
            targets[f] = targets.get(f, []) + [t]
        path, stack = [], ['JFK']
        while stack:
            while targets.get(stack[-1]):
                stack.append(targets[stack[-1]].pop())
            path.append(stack.pop())
        '''
        self.targets = defaultdict(list)
        for f, t in sorted(tickets)[::-1]:
            self.targets[f].append(t)
        path = []
        self.dfs('JFK', path)
        '''
        return path[::-1]

    def dfs(self, airport, path):
        while self.targets[airport]:
            self.dfs(self.targets[airport].pop(), path)
        path.append(airport)

