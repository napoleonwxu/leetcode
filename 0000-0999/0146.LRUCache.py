from collections import OrderedDict

class LRUCache(object):
    def __init__(self, capacity):
        """
        :type capacity: int
        """
        self.cap = capacity
        self.dic = OrderedDict()

    def get(self, key):
        """
        :rtype: int
        """
        if key in self.dic:
            val = self.dic.pop(key)
            self.dic[key] = val
            return val
        return -1

    def set(self, key, value):
        """
        :type key: int
        :type value: int
        :rtype: nothing
        """
        if key in self.dic:
            self.dic.pop(key)
        elif len(self.dic) == self.cap:
            self.dic.popitem(last = False)
        self.dic[key] = value
