class LFUCache:

    def __init__(self, capacity: int):
        self.cache={}
        self.counter={}
        self.capacity=capacity
        self.used=0


    def get(self, key: int) -> int:
        if self.capacity==0:
            return -1
        if self.cache.get(key) is not None:
            self.counter[key]+=1
            return self.cache[key]
        else:
            return -1


    def put(self, key: int, value: int) -> None:
        if self.capacity==0:
            return
        if self.cache.get(key) is not None:
            self.counter[key]+=1
            self.cache[key]=value
        else:
            if self.used<self.capacity:
                self.cache[key]=value
                self.counter[key]=0
                self.used+=1
            else:
                keys=list(self.counter.keys())
                x=lambda x:self.counter[x]
                keys.sort(key=x)
                # 淘汰值最小的
                del self.cache[keys[0]]
                del self.counter[keys[0]]
                self.cache[key]=value
                self.counter[key]=0




# Your LFUCache object will be instantiated and called as such:
# obj = LFUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)