from heapq import heappush, heappop

class Solution(object):
    def getSkyline(self, buildings):
        """
        :type buildings: List[List[int]]
        :rtype: List[List[int]]
        """
        ans = []
        i, length = 0, len(buildings)
        liveHR = []

        while i < length or liveHR:
            if not liveHR or (i < length and buildings[i][0] <= -liveHR[0][1]):
                x = buildings[i][0]
                while i < length and x == buildings[i][0]:
                    heappush(liveHR, (-buildings[i][2], -buildings[i][1]))
                    i += 1
            else:
                x = -liveHR[0][1]
                while liveHR and -liveHR[0][1] <= x:
                    heappop(liveHR)
            height = len(liveHR) and -liveHR[0][0]
            if not ans or height != ans[-1][1]:
                ans.append([x, height])
        return ans
