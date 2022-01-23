# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def isValidBST(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        '''
        if not root:
            return True
        '''
        return self.dfs(root, float('-inf'), float('inf'))
        
    def dfs(self, node, small, big):
        if not node:
            return True
        if node.val <= small or node.val >= big:
            return False
        return self.dfs(node.left, small, node.val) and self.dfs(node.right, node.val, big)
        