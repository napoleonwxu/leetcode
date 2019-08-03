class RandomizedSet(object):

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.nums = []
        self.length = 0
        self.locs = {}

    def insert(self, val):
        """
        Inserts a value to the set. Returns true if the set did not already contain the specified element.
        :type val: int
        :rtype: bool
        """
        if val in self.locs:
            return False
        self.nums.append(val)
        self.locs[val] = self.length
        self.length += 1
        return True

    def remove(self, val):
        """
        Removes a value from the set. Returns true if the set contained the specified element.
        :type val: int
        :rtype: bool
        """
        if val in self.locs:
            loc = self.locs[val]
            last = self.nums[-1]
            self.nums[loc] = last
            self.locs[last] = loc
            self.nums.pop()
            self.locs.pop(val)
            self.length -= 1
            return True
        return False

    def getRandom(self):
        """
        Get a random element from the set.
        :rtype: int
        """
        if self.length > 0:
            return self.nums[random.randint(0, self.length-1)]


# Your RandomizedSet object will be instantiated and called as such:
# obj = RandomizedSet()
# param_1 = obj.insert(val)
# param_2 = obj.remove(val)
# param_3 = obj.getRandom()