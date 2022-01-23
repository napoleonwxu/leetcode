class Solution(object):
    def originalDigits(self, s):
        """
        :type s: str
        :rtype: str
        """
        count = [0] * 10
        for ch in s:
            if ch == 'z': count[0] += 1
            if ch == 'w': count[2] += 1
            if ch == 'u': count[4] += 1
            if ch == 'x': count[6] += 1
            if ch == 'g': count[8] += 1
            if ch == 'h': count[3] += 1 # 3-8
            if ch == 'f': count[5] += 1 # 5-4
            if ch == 's': count[7] += 1 # 7-6
            if ch == 'o': count[1] += 1 # 1-0-2-4
            if ch == 'i': count[9] += 1 # 9-5-6-8
        count[3] -= count[8]
        count[5] -= count[4]
        count[7] -= count[6]
        count[1] -= (count[0] + count[2] + count[4])
        count[9] -= (count[5] + count[6] + count[8])
        return ''.join(str(i)*n for i, n in enumerate(count))
        