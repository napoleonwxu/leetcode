# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Codec:
    def serialize(self, root):
        """Encodes a tree to a single string.
        :type root: TreeNode
        :rtype: str
        """
        def preorder(root):
            if root:
                vals.append(root.val)
                preorder(root.left)
                preorder(root.right)
        vals = []
        preorder(root)
        return ' '.join(map(str, vals))

    def deserialize(self, data):
        """Decodes your encoded data to tree.
        :type data: str
        :rtype: TreeNode
        """
        def buildTree(preorder, inorder):
            if not inorder:
                return None
            idx = inorder.index(preorder.pop(0))
            root = TreeNode(inorder[idx])
            root.left = buildTree(preorder, inorder[:idx])
            root.right = buildTree(preorder, inorder[idx+1:])
            return root
        preorder = map(int, data.split())
        inorder = sorted(preorder)
        return buildTree(preorder, inorder)

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))