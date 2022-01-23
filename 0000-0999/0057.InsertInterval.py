# Definition for an interval.
class Interval(object):
    def __init__(self, s=0, e=0):
        self.start = s
        self.end = e

class Solution(object):
    def insert(self, intervals, newInterval):
        """
        :type intervals: List[Interval]
        :type newInterval: Interval
        :rtype: List[Interval]
        """
        s, e = newInterval.start, newInterval.end
        left, right = [], []
        for pair in intervals:
            if pair.end < s:
                left.append(pair)
            elif pair.start > e:
                right.append(pair)
            else:
                s = min(s, pair.start)
                e = max(e, pair.end)
        return left + [Interval(s, e)] + right
        '''
        ans = []
        for pair in sorted(intervals+[newInterval], key = lambda x: x.start):
            if ans and pair.start <= ans[-1].end:
                ans[-1].end = max(ans[-1].end, pair.end)
            else:
                ans.append(pair)
        return ans
        '''