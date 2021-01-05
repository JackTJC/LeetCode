
# class LRUCache:
#
#     def __init__(self, capacity: int):
#         self.cache={}
#         self.counter={}
#         self.capacity=capacity
#         self.used=0
#
#
#     def get(self, key: int) -> int:
#         for k in self.counter.keys():
#             self.counter[k]+=1
#         if self.cache.get(key):
#             #被访问则置0
#             self.counter[key]=0
#             return self.cache.get(key)
#         else:
#             return -1
#
#
#
#     def put(self, key: int, value: int) -> None:
#         if self.used<self.capacity:
#             self.cache[key]=value
#             self.counter[key]=0
#             self.used+=1
#         else:
#             #选择计数值最大的淘汰
#             keys=list(self.counter.keys())
#             x=lambda x:self.counter[x]
#             keys.sort(key=x)
#             del self.cache[keys[-1]]
#             del self.counter[keys[-1]]
#             #放入新值
#             self.cache[key]=value
#             self.counter[key]=0

import time


class Data:
    def __init__(self, val):
        self.val = val
        self.time = time.time()

    def getVal(self):
        self.time = time.time()
        return self.val

    def update(self, val):
        self.time = time.time()
        self.val = val


class LRUCache:

    def __init__(self, capacity: int):
        self.cache = dict()
        self.capacity = capacity

    def get(self, key: int) -> int:
        ret = self.cache.get(key)
        if ret:
            return ret.getVal()
        else:
            return -1

    def put(self, key: int, value: int) -> None:
        if self.cache.get(key):
            self.cache[key].update(value)
            return
        if len(self.cache) < self.capacity:
            self.cache[key] = Data(value)
            self.capacity += 1
        else:
            k = list(self.cache.keys())
            sortKey = lambda x: self.cache[x].time
            k.sort(key=sortKey)
            del self.cache[k[0]]
            self.cache[key] = Data(value)


if __name__ == '__main__':
    cache = LRUCache(2)
    cache.put(1, 1)
    cache.put(2, 2)
    cache.get(1)
    cache.put(3, 3)
    cache.get(2)
    cache.put(4, 4)
    cache.get(1)
    cache.get(3)
    cache.get(4)
