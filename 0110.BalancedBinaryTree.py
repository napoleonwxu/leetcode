# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def isBalanced(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        '''
        if not root:# or (not root.left and not root.right):
            return True
        if abs(self.depth(root.left)-self.depth(root.right)) > 1:
            return False
        return self.isBalanced(root.left) and self.isBalanced(root.right)
        '''
        self.balance = True
        self.search(root)
        return self.balance

    def depth(self, node):
        if not node:
            return 0
        return max(self.depth(node.left), self.depth(node.right)) + 1

    def search(self, node):
        if not self.balance or not node:
            return 0
        left = self.search(node.left)
        right = self.search(node.right)
        if abs(left-right) > 1:
            self.balance = False
        return max(left, right) + 1

