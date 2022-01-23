# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def isSameTree(self, p, q):
        """
        :type p: TreeNode
        :type q: TreeNode
        :rtype: bool
        """
        if not p and not q:
            return True
        '''
        if p and q and p.val == q.val:
            return self.isSameTree(p.left, q.left) and self.isSameTree(p.right, q.right)
        return False
        '''
        queue = collections.deque([(p, q)])
        while queue:
            pp, qq = queue.popleft()
            if not pp and not qq:
                continue
            elif pp and qq and pp.val == qq.val:
                queue.append((pp.left, qq.left))
                queue.append((pp.right, qq.right))                
            else:
                return False
        return True
        