class Solution(object):
    def numTilePossibilities(self, tiles):
        """
        :type tiles: str
        :rtype: int
        """
        cnt = 0
        for i in xrange(len(tiles), 0, -1):
            cnt += len(set(list(itertools.permutations(tiles, i))))
        return cnt