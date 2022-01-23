class Solution(object):
    def predictPartyVictory(self, senate):
        """
        :type senate: str
        :rtype: str
        """
        queue = collections.deque()
        cnt, ban = [0, 0], [0, 0]
        for ch in senate:
            cnt[ch=='R'] += 1
            queue.append(1 if ch=='R' else 0)
            
        while all(cnt):
            n = queue.popleft()
            if ban[n]:
                ban[n] -= 1
                cnt[n] -= 1
            else:
                ban[n^1] += 1
                queue.append(n)
        return "Radiant" if cnt[1] else "Dire"
        