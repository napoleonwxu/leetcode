# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def buildTree(self, preorder, inorder):
        """
        :type preorder: List[int]
        :type inorder: List[int]
        :rtype: TreeNode
        """
        ''' #1 Memory Limit Exceeded
        if not preorder or not inorder:
            return None
        i = inorder.index(preorder[0])
        node = TreeNode(preorder[0])
        node.left = self.buildTree(preorder[1:i+1], inorder[:i])
        node.right = self.buildTree(preorder[i+1:], inorder[i+1:])
        return node
        '''
        #2
        if not inorder:
            return None
        i = inorder.index(preorder[0])
        node = TreeNode(preorder.pop(0))
        node.left = self.buildTree(preorder, inorder[:i])
        node.right = self.buildTree(preorder, inorder[i+1:])
        return node
        ''' #3
        self.preorder, self.inorder = preorder, inorder
        return self.dfs(0, len(preorder), 0, len(inorder))
        '''

    def dfs(self, preStart, preEnd, inStart, inEnd):
        if preStart >= preEnd or inStart >= inEnd:
            return None
        i = self.inorder[inStart:inEnd].index(self.preorder[preStart])
        node = TreeNode(self.preorder[preStart])
        node.left = self.dfs(preStart+1, preStart+i+1, inStart, inStart+i)
        node.right = self.dfs(preStart+i+1, preEnd, inStart+i+1, inEnd)
        return node
