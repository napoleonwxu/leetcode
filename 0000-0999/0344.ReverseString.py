class Solution(object):
    def reverseString(self, s):
        """
        :type s: str
        :rtype: str
        """
        #return s[::-1]
        return ''.join(reversed(list(s)))

''' #
char* reverseString(char* s) {
    if (s == NULL)
        return s;
    char *begin = s;
    char *end = s;
    while (*end != '\0')
        end ++;
    end --;
    while (begin < end){
        char ch = *begin;
        *begin = *end;
        *end = ch;
        begin ++, end --;
    }
    return s;
}
'''