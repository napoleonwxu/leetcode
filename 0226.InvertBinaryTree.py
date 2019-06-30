# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def invertTree(self, root):
        """
        :type root: TreeNode
        :rtype: TreeNode
        """
        # recursive
        if not root:
            return None
        '''
        l = self.invertTree(root.left)
        r = self.invertTree(root.right)
        root.left = r
        root.right = l
        '''
        root.left, root.right = self.invertTree(root.right), self.invertTree(root.left)
        return root
        """
        # BFS/iteratively with queue
        queue = [root]
        while queue:
            node = queue.pop(0)
            if node:
                node.left, node.right = node.right, node.left
                queue += [node.left, node.right]
        return root
        """