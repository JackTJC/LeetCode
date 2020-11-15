#1338
import collections
from typing import List
class Solution:
    def minSetSize(self, arr: List[int]) -> int:
        if len(arr)==0:
            return -1
        if len(arr)==1:
            return 1
        deleteLen=0
        deleteCount=0
        counter=collections.Counter(arr)
        sortedCounter=counter.most_common(len(counter))
        for k,v in sortedCounter:
            if deleteLen<int(len(arr)/2):
                deleteLen+=v
                deleteCount+=1
            else:
                break
        return deleteCount