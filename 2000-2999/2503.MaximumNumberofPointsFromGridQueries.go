class Solution:
    def maxPoints(self, grid: List[List[int]], queries: List[int]) -> List[int]:
        m, n = len(grid), len(grid[0])
        heap = [(grid[0][0], 0, 0)]
        visited = {(0, 0)}
        order = []
        while len(heap) > 0:
            cur, i, j = heapq.heappop(heap)
            order.append(cur)
            for i, j in [(i-1, j), (i+1, j), (i, j-1), (i, j+1)]:
                if 0 <= i < m and 0 <= j < n and (i, j) not in visited:
                    visited.add((i, j))
                    heapq,heappush(heap, (grid[i][j], i, j))
        tmp = 0
        for i in range(len(order)):
            tmp = max(tmp, order[i])
            order[i] = tmp
        ans = [0] * len(queries)
        for i, q in enumerate(queries):
            ans[i] = bisect.bisect_left(order, q)
        return ans
