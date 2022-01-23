class Solution:
    def smallestChair(self, times: List[List[int]], targetFriend: int) -> int:
        arrivals, leaves = [], []
        for idx, time in enumerate(times):
            heapq.heappush(arrivals, (time[0], idx))
            heapq.heappush(leaves, (time[1], idx))
        occupied = [False] * len(times)
        dic = {}
        while arrivals and leaves:
            if arrivals[0][0] < leaves[0][0]:
                _, idx = heapq.heappop(arrivals)
                chair = occupied.index(False)
                dic[idx] = chair
                occupied[chair] = True
                if idx == targetFriend:
                    return chair
            else:
                _, idx = heapq.heappop(leaves)
                occupied[dic[idx]] = False
                