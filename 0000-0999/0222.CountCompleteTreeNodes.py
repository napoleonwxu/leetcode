# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def countNodes(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        if not root:
            return 0
        l = self.height(root.left)
        r = self.height(root.right)
        if l == r:
            return 2**l + self.countNodes(root.right)
        else:
            return 2**r + self.countNodes(root.left)

    def height(self, node):
        if not node:
            return 0
        return 1 + self.height(node.left)
