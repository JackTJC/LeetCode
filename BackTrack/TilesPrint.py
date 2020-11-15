#1079
import collections
class Solution:
    def numTilePossibilities(self, tiles: str) -> int:
        def backTrack(countDict,res):
            res[0]+=1
            for key in countDict.keys():
                if countDict[key]==0:
                    continue
                countDict[key]-=1
                backTrack(countDict,res)
                countDict[key]+=1
        counter=collections.Counter(tiles)
        countDict=dict(counter)
        res=[0]
        backTrack(countDict,res)
        return res[0]-1