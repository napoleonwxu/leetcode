# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def lowestCommonAncestor(self, root, p, q):
        """
        :type root: TreeNode
        :type p: TreeNode
        :type q: TreeNode
        :rtype: TreeNode
        """
        if not root or root == p or root == q:
            return root
        l = self.lowestCommonAncestor(root.left, p, q)
        r = self.lowestCommonAncestor(root.right, p, q)
        if l and r:
            return root
        return l or r
        '''
        pathp = self.findPath(root, p)
        pathq = self.findPath(root, q)
        ans = None
        i = 0
        while i < min(len(pathp), len(pathq)) and pathp[i] == pathq[i]:
            ans = pathp[i]
            i += 1
        return ans
        '''
        
    def findPath(self, root, node):
        path = []
        last = None
        while path or root:
            if root:
                path.append(root)
                root = root.left
            else:
                temp = path[-1]
                if temp.right and temp != last:
                    root = temp.right
                else:
                    if temp == node:
                        return path
                    last = path.pop()
                    root = None
        return path
        
