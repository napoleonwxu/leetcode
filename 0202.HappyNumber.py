class Solution(object):
    def isHappy(self, n):
        """
        :type n: int
        :rtype: bool
        """
        if n <= 0:
            return False

        num = n
        res = []
        while True:
            add = 0
            while num > 0:
                add += (num%10) ** 2
                num = num / 10
            if add == 1:
                return True
            else:
                for i in range(len(res)):
                    if add == res[i]:
                        return False
            res.append(add)
            num = add
            print res

test = Solution()
print test.isHappy(3)
'''
for i in range(15):
    print i, test.isHappy(i)
'''