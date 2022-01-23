class Solution(object):
    def findCheapestPrice(self, n, flights, src, dst, K):
        """
        :type n: int
        :type flights: List[List[int]]
        :type src: int
        :type dst: int
        :type K: int
        :rtype: int
        """
        # Dijkstra + Heap
        f = collections.defaultdict(dict)
        for s, d, p in flights:
            f[s][d] = p
        heap = [(0, src, K+1)]
        while heap:
            price, city, fly = heapq.heappop(heap)
            if city == dst:
                return price
            if fly > 0:
                for nxt in f[city]:
                    heapq.heappush(heap, (price+f[city][nxt], nxt, fly-1))
        return -1
