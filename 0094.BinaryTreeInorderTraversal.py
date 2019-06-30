# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def inorderTraversal(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        # Iteratively with stack
        ans = []
        stack = []
        node = root
        while stack or node:
            if node:
                stack.append(node)
                node = node.left
            else:
                node = stack.pop()
                ans.append(node.val)
                node = node.right
        return ans
        '''
        # Recursive
        if not root:
            return []
        #return self.inorderTraversal(root.left) + [root.val] + self.inorderTraversal(root.right)
        ans = []
        ans += self.inorderTraversal(root.left)
        ans.append(root.val)
        ans += self.inorderTraversal(root.right)
        return ans
        '''