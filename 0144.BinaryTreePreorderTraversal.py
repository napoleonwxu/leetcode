# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def preorderTraversal(self, root):
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
                ans.append(node.val)
                node = node.left
            else:
                node = stack.pop()
                node = node.right
        return ans
        '''
        # Recursive
        if not root:
            return []
        #return [root.val] + self.preorderTraversal(root.left) + self.preorderTraversal(root.right)
        ans = []
        ans.append(root.val)
        ans += self.preorderTraversal(root.left)
        ans += self.preorderTraversal(root.right)
        return ans
        '''
