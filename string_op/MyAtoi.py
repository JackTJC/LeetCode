# 8. 字符串转换整数 (atoi)

class Solution:
    def myAtoi(self, s: str) -> int:
        if len(s) == 0:
            return 0
        i = 0
        # TODO 去除空格
        while i <= len(s) - 1 and s[i] == ' ':
            i += 1

        # TODO 获得符号
        sign = 1  # 1 表示正数 0表示负数
        if i <= len(s) - 1 and (s[i] == '+' or s[i] == '-'):
            if s[i] == '-':
                sign = 0
            i += 1
        if i == len(s) or not s[i].isdigit():
            return 0
        # TODO 求得数值
        sum = 0
        while i <= len(s) - 1 and s[i].isdigit():
            sum *= 10
            sum += int(s[i])
            i += 1
        sum = sum if sign == 1 else -sum

        # TODO 溢出判断
        if sum < -2 ** 31:
            return -2 ** 31
        if sum > 2 ** 31 - 1:
            return 2 ** 31 - 1
        return sum
