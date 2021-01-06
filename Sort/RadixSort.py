from typing import List
import random


# 2**31-1=2147483647
# 共10位
def radixSort(arr: List[int]):
    base = 1  # 为什么是base=1？ 因为个位数的获得来自公式 (number // 1)%10
    for i in range(10):
        buckets = [[] for _ in range(10)] # 每次迭代都要重置
        for num in arr:
            buckets[(num // base) % 10].append(num)
        arr = []  # 每次迭代都要重置
        for bucket in buckets:
            for num in bucket:
                arr.append(num)
        base *= 10
    return arr


if __name__ == '__main__':
    print(radixSort([random.randint(1, 10000) for _ in range(20)]))
