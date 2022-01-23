class Solution:
    def maxLength(self, arr: List[str]) -> int:
        unique = [set()]
        for a in arr:
            s = set(a)
            if len(s) != len(a):
                continue
            for t in unique[:]:
                if not s&t:
                    unique.append(s|t)
        return max(len(s) for s in unique)