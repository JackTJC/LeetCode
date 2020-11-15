# 快手 2020/11/12 后端实习
def parseInt(s: str):
    i = 1
    sum = 0
    flag = True
    for c in s[::-1]:
        if (c.isdigit() or c == '-' or c == '+') and flag:
            if c.isdigit():
                sum += (ord(c) - ord('0')) * i
                i *= 10
            elif c == '-':
                sum = -sum
                flag = False
            else:
                flag = False
        else:
            raise ValueError
    return sum


if __name__ == '__main__':
    a = parseInt("+7656342341345457897067564235786790890954359089134256890")
    print(a)
