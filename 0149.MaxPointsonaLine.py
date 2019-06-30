# Definition for a point.
class Point(object):
    def __init__(self, a=0, b=0):
        self.x = a
        self.y = b

class Solution(object):
    def maxPoints(self, points):
        """
        :type points: List[Point]
        :rtype: int
        """
        length = len(points)
        if length <= 2:
            return length
        ans = 0
        for i in xrange(length-1):
            dic = {'i': 1}
            same = 0
            xi, yi = points[i].x, points[i].y
            for j in xrange(i+1, length):
                xj, yj = points[j].x, points[j].y
                if xi == xj and yi == yj:
                    same += 1
                    continue
                if xi == xj:
                    scope = 'i'
                else:
                    scope = float(yi-yj)/(xi-xj)
                if scope not in dic:
                    dic[scope] = 1
                dic[scope] += 1
            ans = max(ans, max(dic.values())+same)
        return ans

print Solution().maxPoints([Point(0,0), Point(1,1), Point(2,2), Point(1,-1)])