class Solution(object):
    def getHint(self, secret, guess):
        """
        :type secret: str
        :type guess: str
        :rtype: str
        """
        a = b = 0
        copy_s = []
        copy_g = []
        for i in xrange(len(guess)):
            if guess[i] == secret[i]:
                a += 1
            else:
                copy_s.append(secret[i])
                copy_g.append(guess[i])
        for ch in copy_g:
            if ch in copy_s:
                b += 1
                copy_s.remove(ch)
        return str(a) + 'A' + str(b) + 'B'