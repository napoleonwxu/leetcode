# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def rightSideView(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        if not root:
            return []
        '''
        ans = []
        currLine = [root]
        nextLine = []
        while currLine:
            ans.append(currLine[-1].val)
            while currLine:
                node = currLine.pop(0)
                if node.left:
                    nextLine.append(node.left)
                if node.right:
                    nextLine.append(node.right)
            currLine, nextLine = nextLine, currLine
        return ans
        '''
        r = self.rightSideView(root.right)
        l = self.rightSideView(root.left)
        return [root.val] + r + l[len(r):]
        