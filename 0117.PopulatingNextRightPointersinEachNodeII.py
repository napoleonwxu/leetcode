# Definition for binary tree with next pointer.
# class TreeLinkNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
#         self.next = None

class Solution(object):
    def connect(self, root):
        """
        :type root: TreeLinkNode
        :rtype: nothing
        """
        if not root:
            return
        #currLine = collections.deque([root])
        #nextLine = collections.deque()
        currLine, nextLine = [root], []
        while currLine:
            #node = currLine.popleft()
            node = currLine.pop(0)
            if node.left:
                nextLine.append(node.left)
            if node.right:
                nextLine.append(node.right)
            if currLine:
                node.next = currLine[0]
            if not currLine and nextLine:
                #currLine, nextLine = nextLine, currLine
                currLine, nextLine = nextLine, []
