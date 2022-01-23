# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Codec:
    def serialize(self, root):
        """Encodes a tree to a single string.
        :type root: TreeNode
        :rtype: str
        """
        ans = ''
        queue = [root]
        while queue:
            node = queue.pop(0)
            ans += '#'
            if node:
                ans += str(node.val)
                queue += [node.left, node.right]
        return ans

    def deserialize(self, data):
        """Decodes your encoded data to tree.
        :type data: str
        :rtype: TreeNode
        """
        if data == '#':
            return None
        vals = data.lstrip('#').split('#')
        root = TreeNode(int(vals[0]))
        queue = [root]
        i = 1
        while queue:# and i < len(vals):
            node = queue.pop(0)
            if vals[i]:
                node.left = TreeNode(int(vals[i]))
                queue.append(node.left)
            i += 1
            if vals[i]:
                node.right = TreeNode(int(vals[i]))
                queue.append(node.right)
            i += 1
        return root

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))
# Your Codec object will be instantiated and called as such:
# [5,2,3,null,null,2,4,3,1]
root = TreeNode(5)
root.left = TreeNode(2)
n = root.right = TreeNode(3)
n.left = TreeNode(2)
n.right = TreeNode(4)
n = n.left
n.left = TreeNode(3)
n.right = TreeNode(1)

codec = Codec()
s = codec.serialize(root)
print s
print codec.deserialize(s)