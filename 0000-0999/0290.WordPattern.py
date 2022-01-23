class Solution(object):
    def wordPattern(self, pattern, str):
        """
        :type pattern: str
        :type str: str
        :rtype: bool
        """
        str_list = str.split(' ')
        print str_list
        dict = {}
        if len(pattern) == len(str_list):
            for i in xrange(len(pattern)):
                #print pattern[i]
                print dict
                if pattern[i] not in dict.keys() and str_list[i] not in dict.values():
                    dict[pattern[i]] = str_list[i]
                elif pattern[i] not in dict.keys() or str_list[i] not in dict.values():
                    return False
                elif dict[pattern[i]] != str_list[i]:
                    return False
        else:
            return False
        return True

test = Solution()
res = test.wordPattern('abba', 'dog dog dog dog')
print res
