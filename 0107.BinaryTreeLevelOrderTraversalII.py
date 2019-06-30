# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def levelOrderBottom(self, root):
        """
        :type root: TreeNode
        :rtype: List[List[int]]
        """
        if not root:
            return []
        ans = []
        level = [root]
        while level:
            #ans.append([node.val for node in level])
            ans.insert(0, [node.val for node in level])
            level = [child for f in level for child in (f.left, f.right) if child]
        #ans.reverse()
        return ans
        '''
        ans = []
        while temp:
            ans.append(temp.pop())
        return ans
        '''