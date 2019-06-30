class Queue(object):
    def __init__(self):
        """
        initialize your data structure here.
        """
        self.stackIn, self.stackOut = [], []

    def push(self, x):
        """
        :type x: int
        :rtype: nothing
        """
        self.stackIn.append(x)

    def pop(self):
        """
        :rtype: nothing
        """
        self.update()
        if self.stackOut:
            self.stackOut.pop()

    def peek(self):
        """
        :rtype: int
        """
        self.update()
        if self.stackOut:
            return self.stackOut[-1]

    def empty(self):
        """
        :rtype: bool
        """
        return not self.stackIn and not self.stackOut

    def update(self):
        if not self.stackOut:
            while self.stackIn:
                self.stackOut.append(self.stackIn.pop())
