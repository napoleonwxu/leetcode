class Solution(object):
    def numberToWords(self, num):
        """
        :type num: int
        :rtype: str
        """
        if num == 0:
            return 'Zero'
        d = ['', ' Thousand', ' Million', ' Billion']
        ans = ''
        for f in xrange(4):
            if num == 0:
                break
            t = num % 1000
            if t > 0:
                if ans:
                    ans = ' ' + ans
                ans = self.helper(t) + d[f] + ans
            num /= 1000            
        '''
        f = 0
        while num > 0:
            t = num % 1000
            if t > 0:
                if ans:
                    ans = ' ' + ans
                ans = self.helper(t) + d[f] + ans
            num /= 1000
            f += 1
        '''
        return ans
        
    def helper(self, num):
        if num == 0:
            return ''
        d1 = ['', 'One', 'Two', 'Three', 'Four', 'Five', 'Six', 'Seven', 'Eight', 'Nine']
        d2 = ['Ten', 'Eleven', 'Twelve', 'Thirteen', 'Fourteen', 'Fifteen', 'Sixteen', 'Seventeen', 'Eighteen', 'Nineteen']
        d3 = ['', '', 'Twenty', 'Thirty', 'Forty', 'Fifty', 'Sixty', 'Seventy', 'Eighty', 'Ninety']
        d4 = [' Hundred']
        ans = ''
        t = num/100
        if t > 0:
            ans += d1[t] + d4[0]
            num %= 100
        t = num/10
        if t == 1:
            if ans:
                ans += ' '
            ans += d2[num%10]
            return ans
        elif t >= 2:
            if ans:
                ans += ' '
            ans += d3[t]
            num %= 10
        if num > 0:
            if ans:
                ans += ' '
            ans += d1[num]
        return ans
        