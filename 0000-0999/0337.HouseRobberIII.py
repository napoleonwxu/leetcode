# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def rob(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        return self.dfs(root)[1]

    def dfs(self, node):
        if not node:
            return (0, 0)
        l0, l1 = self.dfs(node.left)
        r0, r1 = self.dfs(node.right)
        return l1+r1, max(l1+r1, l0+r0+node.val)
