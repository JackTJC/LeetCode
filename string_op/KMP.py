def getNext(s):
    Next = []
    x = 1
    now = 0
    Next.append(0)
    while x < len(s):
        if s[now] == s[x]:
            now += 1
            x += 1
            Next.append(now)
        elif now:
            now = Next[now - 1]
        else:
            Next.append(0)
            x += 1
    return Next


if __name__ == '__main__':
    s = "abcabdddabcabc"
    print(getNext(s))
    print(list(s))
