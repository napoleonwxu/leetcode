# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def recoverTree(self, root):
        """
        :type root: TreeNode
        :rtype: void Do not return anything, modify root in-place instead.
        """
        self.first = self.second = None
        self.pre = TreeNode(float('-inf'))
        self.traverse(root)
        self.first.val, self.second.val = self.second.val, self.first.val

    def traverse(self, root):
        if not root:
            return
        self.traverse(root.left)
        if not self.first and self.pre.val >= root.val:
            self.first = self.pre
        if self.first and self.pre.val >= root.val:
            self.second = root
        self.pre = root
        self.traverse(root.right)
