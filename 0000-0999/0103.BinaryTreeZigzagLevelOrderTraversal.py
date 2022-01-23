# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def zigzagLevelOrder(self, root):
        """
        :type root: TreeNode
        :rtype: List[List[int]]
        """
        if not root:
            return []
        level = [root]
        ans = []
        flag = True
        while level:
            temp = [node.val for node in level]
            level = [child for f in level for child in (f.left, f.right) if child]
            if not flag:
                temp.reverse()
            ans.append(temp)
            flag = not flag
        return ans
