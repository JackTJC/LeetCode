
class LRUCache:

    def __init__(self, capacity: int):
        self.cache={}
        self.counter={}
        self.capacity=capacity
        self.used=0


    def get(self, key: int) -> int:
        for k in self.counter.keys():
            self.counter[k]+=1
        if self.cache.get(key):
            #被访问则置0
            self.counter[key]=0
            return self.cache.get(key)
        else:
            return -1



    def put(self, key: int, value: int) -> None:
        if self.used<self.capacity:
            self.cache[key]=value
            self.counter[key]=0
            self.used+=1
        else:
            #选择计数值最大的淘汰
            keys=list(self.counter.keys())
            x=lambda x:self.counter[x]
            keys.sort(key=x)
            del self.cache[keys[-1]]
            del self.counter[keys[-1]]
            #放入新值
            self.cache[key]=value
            self.counter[key]=0




if __name__ == '__main__':
    cache=LRUCache(2)

    cache.get(2)
    cache.put(2,6)
    cache.get(1)
    cache.put(1,5)
    cache.put(1,2)
    cache.get(1)
    cache.get(2)
