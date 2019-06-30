class Solution(object):
    def reverseVowels(self, s):
        """
        :type s: str
        :rtype: str
        """
        t = list(s)
        vowels = set(list('aeiouAEIOU'))
        i, j = 0, len(s) - 1
        while i < j:
            while i < j and t[i] not in vowels:
                i += 1
            while j > i and t[j] not in vowels:
                j -= 1
            t[i], t[j] = t[j], t[i]
            i += 1
            j -= 1
        return ''.join(t)

'''
int isVowel(char ch) {
    if (ch=='a' || ch=='e' || ch=='i' || ch=='o' || ch=='u' || ch=='A' || ch=='E' || ch=='I' || ch=='O' || ch=='U')
        return 1;
    return 0;
}

char* reverseVowels(char* s) {
    if (s == NULL)
        return s;
    char *start = s;
    char *end = s;
    while (*end != '\0')
        end ++;
    end --;
    while (start < end) {
        while (start < end && isVowel(*start) == 0)
            start ++;
        while (end > start && isVowel(*end) == 0)
            end --;
        char ch = *start;
        *start = *end;
        *end = ch;
        start ++, end --;
    }
    return s;
}
'''