class Solution(object):
    def combinationSum2(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        ans = []
        self.dfs(sorted(candidates), target, 0, [], ans)
        return ans

    def dfs(self, cand, target, index, temp, ans):
        if target == 0:
            ans.append(temp)
            return
        for i in xrange(index, len(cand)):
            if target < cand[i]:
                return
            if i > index and cand[i] == cand[i-1]:
                continue
            self.dfs(cand, target-cand[i], i+1, temp+[cand[i]], ans)