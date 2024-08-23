import threading
import time

class TokenBucket:
    def __init__(self, rate, capacity):
        self.rate = rate
        self.capacity = capacity
        self.tokens = capacity
        self.lastCheck = time.time()
        self.lock = threading.Lock()

    def  refillToken(self):
        currTime = time.time()
        gap = currTime - self.lastCheck
        lastCheck = currTime

        # it should be minium of the maxium token that can be avail to user and the refillTokens (gap * rate)
        self.tokens = min(self.capacity, self.tokens + gap * self.rate)
        print(self.tokens)

    def consumeToken(self):
        with self.lock:
            self.refillToken()    
            if self.tokens > 1:
                self.tokens -= 1
                return True
            else:
                return False    
tokenBucketObj = TokenBucket(5, 10)
for i in range(20):
    if tokenBucketObj.consumeToken():
        print("Request succeeded")
    else:
        print("Request Failed")    