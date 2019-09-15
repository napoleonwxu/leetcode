class Codec:
    def __init__(self):
        self.char = string.ascii_letters + string.digits
        self.long2short = {}
        self.short2long = {}

    def encode(self, longUrl):
        """Encodes a URL to a shortened URL.
        
        :type longUrl: str
        :rtype: str
        """
        while longUrl not in self.long2short:
            shortUrl = ""
            for _ in xrange(6):
                shortUrl += random.choice(self.char)
            if shortUrl not in self.short2long:
                self.long2short[longUrl] = shortUrl
                self.short2long[shortUrl] = longUrl
        return self.long2short[longUrl]
        

    def decode(self, shortUrl):
        """Decodes a shortened URL to its original URL.
        
        :type shortUrl: str
        :rtype: str
        """
        if shortUrl in self.short2long:
            return self.short2long[shortUrl]
        return "invalid"
        

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.decode(codec.encode(url))