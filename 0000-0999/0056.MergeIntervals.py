# Definition for an interval.
# class Interval(object):
#     def __init__(self, s=0, e=0):
#         self.start = s
#         self.end = e

class Solution(object):
    def merge(self, intervals):
        """
        :type intervals: List[Interval]
        :rtype: List[Interval]
        """
        ans = []
        for pair in sorted(intervals, key = lambda x: x.start):
            if ans and pair.start <= ans[-1].end:
                ans[-1].end = max(pair.end, ans[-1].end)
            else:
                ans.append(pair)
                #ans.extend([pair])
        return ans