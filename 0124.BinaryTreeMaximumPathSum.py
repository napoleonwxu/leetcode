# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def __init__(self):
        self.maxVal = float('-inf')

    def maxPathSum(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        self.dfs(root)
        return self.maxVal

    def dfs(self, node):
        if not node:
            return 0
        left = max(0, self.dfs(node.left))
        right = max(0, self.dfs(node.right))
        self.maxVal = max(self.maxVal, left+node.val+right)
        return max(left, right) + node.val
