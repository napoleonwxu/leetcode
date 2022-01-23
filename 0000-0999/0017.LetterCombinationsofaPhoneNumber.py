class Solution(object):
    def letterCombinations(self, digits):
        """
        :type digits: str
        :rtype: List[str]
        """
        if not digits:
            return []
        dic = {'1':None, '2':'abc', '3':'def', '4':'ghi', '5':'jkl', '6':'mno', '7':'pqrs', '8':'tuv', '9':'wxyz', '0':' '}
        ans = []
        self.combin(dic, digits, 0, '', ans)
        return ans

    def combin(self, dic, digits, i, comb, ans):
        if i == len(digits):
            ans.append(comb)
            return
        for ch in dic[digits[i]]:
            self.combin(dic, digits, i+1, comb+ch, ans)