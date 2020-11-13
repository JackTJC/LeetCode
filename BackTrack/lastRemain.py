#390 time exceed
class Solution:
    def lastRemaining(self, n: int) -> int:
        return delNumber(list(range(1, n + 1)), True)


def delNumber(numberList, flag):
    if len(numberList) == 1:
        return numberList[0]
    if flag:
        return delNumber([numberList[i] for i in range(1, len(numberList), 2)], not flag)
    else:
        return delNumber([numberList[i] for i in range(len(numberList)-2, -1, -2)], not flag)



if __name__ == '__main__':
    Solution.lastRemaining(None, 11)
