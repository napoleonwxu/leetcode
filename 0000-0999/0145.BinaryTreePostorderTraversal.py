# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def postorderTraversal(self, root):
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
                node = node.right
            else:
                node = stack.pop()
                node = node.left
        return ans[::-1]
        '''
        # Recursive
        if not root:
            return []
        #return self.postorderTraversal(root.left) + self.postorderTraversal(root.right) + [root.val]
        ans = []
        ans += self.postorderTraversal(root.left)
        ans += self.postorderTraversal(root.right)
        ans.append(root.val)
        return ans
        '''
