# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def pathSum(self, root, sum):
        """
        :type root: TreeNode
        :type sum: int
        :rtype: int
        """
        #fast: O(n)
        return self.dfs2(root, sum, [sum])
        
        #Slow: O(n*logn) - O(n*n)
        if not root:
            return 0
        return self.dfs(root, sum) + self.pathSum(root.left, sum) + self.pathSum(root.right, sum)
    
    def dfs2(self, root, sum, targets):
        if not root:
            return 0
        count = 0
        for t in targets:
            if root.val == t:
                count += 1
        targets = [t-root.val for t in targets] + [sum]
        count += self.dfs2(root.left, sum, targets) + self.dfs2(root.right, sum, targets)
        return count
        
    def dfs(self, root, sum):
        if not root:
            return 0
        count = 0
        if root.val == sum:
            count += 1
        count += self.dfs(root.left, sum-root.val) + self.dfs(root.right, sum-root.val)
        return count
        