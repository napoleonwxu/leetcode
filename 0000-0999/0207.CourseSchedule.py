class Solution(object):
    def canFinish(self, numCourses, prerequisites):
        """
        :type numCourses: int
        :type prerequisites: List[List[int]]
        :rtype: bool
        """
        #graph = [[] for course in xrange(numCourses)]
        graph = {course: [] for course in xrange(numCourses)}
        inDegree = [0] * numCourses
        for pair in prerequisites:
            graph[pair[1]].append(pair[0])
            inDegree[pair[0]] += 1
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
        return len(topo) == numCourses