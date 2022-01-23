class Stack(object):
    def __init__(self):
        """
        initialize your data structure here.
        """
        self.queue = []

    def push(self, x):
        """
        :type x: int
        :rtype: nothing
        """
        self.queue.append(x)
        for i in xrange(len(self.queue)-1):
            self.queue.append(self.queue.pop(0))

    def pop(self):
        """
        :rtype: nothing
        """
        if self.queue:
            self.queue.pop(0)

    def top(self):
        """
        :rtype: int
        """
        if self.queue:
            return self.queue[0]

    def empty(self):
        """
        :rtype: bool
        """
        return not self.queue
