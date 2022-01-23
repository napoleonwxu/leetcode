import fractions

class Solution(object):
    def hasGroupsSizeX(self, deck):
        """
        :type deck: List[int]
        :rtype: bool
        """
        return reduce(fractions.gcd, collections.Counter(deck).values()) > 1
        