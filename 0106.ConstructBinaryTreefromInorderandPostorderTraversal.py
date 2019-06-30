# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def buildTree(self, inorder, postorder):
        """
        :type inorder: List[int]
        :type postorder: List[int]
        :rtype: TreeNode
        """
        ''' #1 MLE
        if not inorder or not postorder:
            return None
        i = inorder.index(postorder[-1])
        node = TreeNode(postorder[-1])
        node.left = self.buildTree(inorder[:i], postorder[:i])
        node.right = self.buildTree(inorder[i+1:], postorder[i:-1])
        return node
        '''
        #2
        if not inorder:
            return None
        i = inorder.index(postorder[-1])
        node = TreeNode(postorder.pop())
        node.right = self.buildTree(inorder[i+1:], postorder)
        node.left = self.buildTree(inorder[:i], postorder)
        return node
        ''' #3
        self.inorder, self.postorder = inorder, postorder
        return self.dfs(0, len(inorder), 0, len(postorder))
        '''

    def dfs(self, inStart, inEnd, postStart, postEnd):
        if inStart >= inEnd or postStart >= postEnd:
            return None
        i = self.inorder[inStart:inEnd].index(self.postorder[postEnd-1])
        node = TreeNode(self.postorder[postEnd-1])
        node.left = self.dfs(inStart, inStart+i, postStart, postStart+i)
        node.right = self.dfs(inStart+i+1, inEnd, postStart+i, postEnd-1)
        return node
