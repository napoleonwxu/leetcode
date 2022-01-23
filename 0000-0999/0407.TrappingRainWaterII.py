class Solution(object):
    def trapRainWater(self, heightMap):
        """
        :type heightMap: List[List[int]]
        :rtype: int
        """
        if not heightMap or not heightMap[0]:
            return 0
        m, n = len(heightMap), len(heightMap[0])
        ans = 0
        heap = []
        visited = [[False]*n for i in xrange(m)]
        for i in [0, m-1]:
            for j in xrange(n):
                heapq.heappush(heap, (heightMap[i][j], i, j))
                visited[i][j] = True
        for j in [0, n-1]:
            for i in xrange(1, m-1):
                heapq.heappush(heap, (heightMap[i][j], i, j))
                visited[i][j] = True
        while heap:
            height, x, y = heapq.heappop(heap)
            for i, j in [(x-1, y), (x+1, y), (x, y-1), (x, y+1)]:
                if 0 <= i < m and 0 <= j < n and not visited[i][j]:
                    ans += max(0, height-heightMap[i][j])
                    heapq.heappush(heap, (max(height, heightMap[i][j]), i, j))
                    visited[i][j] = True
        return ans
        