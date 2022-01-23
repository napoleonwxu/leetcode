class MinStack(object):
    def __init__(self):
        """
        initialize your data structure here.
        """
        self.data = []
        self.min = []

    def push(self, x):
        """
        :type x: int
        :rtype: nothing
        """
        '''
        min_x = self.getMin()
        if min_x == None or x < min_x:
            min_x = x 
        self.data.append((x, min_x))
        '''
        self.data.append(x)
        if x <= self.getMin() or self.getMin() is None:
            self.min.append(x)

    def pop(self):
        """
        :rtype: nothing
        """
        '''
        if self.data:
            self.data.pop()
        '''
        if self.data[-1] == self.min[-1]:
            self.min.pop()
        self.data.pop()        
        
    def top(self):
        """
        :rtype: int
        """
        '''
        if self.data:
            return self.data[-1][0]
        else:
            return None
        '''
        if self.data:
            return self.data[-1]
            
    def getMin(self):
        """
        :rtype: int
        """
        '''
        if self.data:
            return self.data[-1][1]
        else:
            return None
        '''
        if self.min:
            return self.min[-1]
        else:
            return None