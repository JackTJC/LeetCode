from typing import List


class Solution:
    def minSumOfLengths(self, arr: List[int], target: int) -> int:
        lengthList = list()
        start = 0
        end = 0
        sum = arr[0]
        while end < len(arr) - 1:
            if sum == target:
                lengthList.append(end - start + 1)
                start = end + 1
                end += 1
                sum = 0
            elif sum < target:
                end += 1
                sum += arr[end]
            else:
                sum -= arr[start]
                end += 1
                start += 1
                sum += arr[end]
        while start <= len(arr) - 1:
            if sum == target:
                lengthList.append(end - start + 1)
            else:
                sum -= arr[start]
                start += 1
        if len(lengthList) < 2:
            return -1
        lengthList.sort()
        return lengthList[0] + lengthList[1]


if __name__ == '__main__':
    Solution.minSumOfLengths(None, [3, 2, 2, 4, 3], 3)
