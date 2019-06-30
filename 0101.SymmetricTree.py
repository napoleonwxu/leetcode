# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def isSymmetric(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        if not root:
            return True
        #return self.equal(root.left, root.right)
        q = [(root.left, root.right)]
        while q:
            left, right = q.pop()
            if not left and not right:
                continue
            elif left and right and left.val == right.val:
                q.append((left.left, right.right))
                q.append((left.right, right.left))
            else:
                return False
        return True


    def equal(self, left, right):
        if not left and not right:
            return True
        if not left or not right:
            return False
        if left.val != right.val:
            return False
        return self.equal(left.left, right.right) and self.equal(left.right, right.left)

