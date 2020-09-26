class Solution:
    def peopleIndexes(self, favoriteCompanies: List[List[str]]) -> List[int]:
        dic = {}
        for i, company in enumerate(favoriteCompanies):
            dic[i] = set(company)
        ans = []
        for i, company in enumerate(favoriteCompanies):
            s = set(company)
            flag = True
            for idx, companies in dic.items():
                if i != idx and s.issubset(companies):
                    flag = False
                    break
            if flag:
                ans.append(i)
        return ans
        