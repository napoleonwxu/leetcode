# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def generateTrees(self, n):
        """
        :type n: int
        :rtype: List[TreeNode]
        """
        if n <= 0:
            return []
        return self.BST(1, n)

    def BST(self, start, end):
        if start > end:
            return [None]
        ans = []
        for val in xrange(start, end+1):
            for l in self.BST(start, val-1):
                for r in self.BST(val+1, end):
                    node = TreeNode(val)
                    node.left = l
                    node.right = r
                    ans += [node]
        return ans
