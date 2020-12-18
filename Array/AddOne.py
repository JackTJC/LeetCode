from typing import List
class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        Cout=0
        res=[]
        for number in digits[::-1]:
            if len(res)==0:
                Si=(number+1)%10
                if (number+1)>=10:
                    Cout=1
                else:
                    Cout=0
            else:
                Si=(number+Cout)%10
                if (number+Cout)>=c10:
                    Cout=1
                else:
                    Cout=0
            res.append(Si)
        if Cout==1:
                res.append(1)
        return res[::-1]

if __name__ == '__main__':
    Solution.plusOne(None,[8,9,9,9])