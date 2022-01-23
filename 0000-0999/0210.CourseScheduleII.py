class Solution(object):
    def findOrder(self, numCourses, prerequisites):
        """
        :type numCourses: int
        :type prerequisites: List[List[int]]
        :rtype: List[int]
        """
        graph = {i:[] for i in xrange(numCourses)}
        inDegree = {i:0 for i in xrange(numCourses)}
        for ch, fa in prerequisites:
            graph[fa].append(ch)
            inDegree[ch] += 1
        topo = []
        for course in xrange(numCourses):
            if inDegree[course] == 0:
                topo.append(course)
        i = 0
        while i < len(topo):
            for course in graph[topo[i]]:
                inDegree[course] -= 1
                if inDegree[course] == 0:
                    topo.append(course)
            i += 1
        if len(topo) == numCourses:
            return topo
        return []