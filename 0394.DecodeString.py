class Solution(object):
    def decodeString(self, s):
        """
        :type s: str
        :rtype: str
        """
        while '[' in s:
            s = re.sub(r'(\d+)\[([a-z]*)\]', lambda m: int(m.group(1)) * m.group(2), s)
        return s
        ans = ''
        stack_num = []
        stack_str = []
        num = 0
        string = ''
        for ch in s:
            if ch.isdigit():
                num = 10*num + int(ch)
            elif ch == '[':
                stack_num.append(num)
                stack_str.append(string)
                num = 0
                string = ''
            elif ch == ']':
                string = stack_str.pop() + stack_num.pop()*string
            else:
                string += ch
        return string
        