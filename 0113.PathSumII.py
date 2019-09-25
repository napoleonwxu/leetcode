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
        :rtype: List[List[int]]
        """
        ans = []
        self.dfs(root, sum, [], ans)
        return ans
        
    def dfs(self, root, sum, path, ans):
        if not root:
            return
        if not root.left and not root.right:
            if sum == root.val:
                ans.append(path+[root.val])
            return
        self.dfs(root.left, sum-root.val, path+[root.val], ans)
        self.dfs(root.right, sum-root.val, path+[root.val], ans)
        