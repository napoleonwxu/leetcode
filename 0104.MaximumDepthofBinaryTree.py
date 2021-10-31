# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def maxDepth(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        if not root:
            return 0
        #return max(self.maxDepth(root.left), self.maxDepth(root.right)) + 1
        '''
        depth = 0
        queue = [(root, 1)]
        while queue:
            node, depth = queue.pop(0)
            queue += [(n, depth+1) for n in [node.left, node.right] if n]
        return depth
        '''
        q = collections.deque([(root, 1)])
        while q:
            node, leval = q.popleft()
            if node.left:
                q.append((node.left, leval+1))
            if node.right:
                q.append((node.right, leval+1))
        return leval
        