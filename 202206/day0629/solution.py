import random


# 535. TinyURL 的加密与解密
class Codec:

    def __init__(self):
        self.urls = {}
        self.tiny_urls = {}

    def encode(self, longUrl: str) -> str:
        if longUrl in self.urls:
            return self.urls[longUrl]
        tiny_url = "http://t.cn/"
        for i in range(5):
            tiny_url += chr(random.randint(ord('a'), ord('z')))
        if tiny_url not in self.tiny_urls:
            self.urls[longUrl] = tiny_url
            self.tiny_urls[tiny_url] = longUrl
            return tiny_url
        return self.encode(longUrl)

    def decode(self, shortUrl: str) -> str:
        return self.tiny_urls[shortUrl]
