# 10. 正则表达式匹配
# 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
#
# '.' 匹配任意单个字符
# '*' 匹配零个或多个前面的那一个元素
# 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        ONLY_ONE = 0
        ANY_SINGLE = 1
        PRE_MULTI = 2
        # TODO 获得模式
        character, mode = Solution.getMode(None,p)
        # TODO 匹配
        i = 0
        j = 0
        while j <= len(s) - 1 and i <= len(character) - 1:
            c = character[i]
            m = mode[i]
            if m == ONLY_ONE:
                if s[j] == c:
                    j += 1
                else:
                    return False
            elif m == ANY_SINGLE:
                if j > len(s) - 1:
                    return False
                else:
                    j += 1
            else:
                if c == '.':
                    while j <= len(s) - 1:
                        j += 1
                else:
                    if s[j]!=c:
                        i+=1
                        continue
                    while j <= len(s) - 1 and s[j] == c:
                        j += 1
            i += 1
        if i<=len(mode)-1:
            while mode[i]==2:
                i+=1
        return i > len(character) - 1 and j > len(s) - 1

    def getMode(self, p):
        ONLY_ONE = 0
        ANY_SINGLE = 1
        PRE_MULTI = 2
        character = []
        mode = []
        i = 0
        # TODO 获得模式
        while i <= len(p) - 1:
            if p[i] == '.':
                character.append('.')
                if i == len(p) - 1:
                    mode.append(ANY_SINGLE)
                else:
                    if p[i + 1] == '*':
                        mode.append(PRE_MULTI)
                        i += 1
                    else:
                        mode.append(ANY_SINGLE)
            else:
                character.append(p[i])
                if i == len(p) - 1:
                    mode.append(ONLY_ONE)
                else:
                    if p[i + 1] == '*':
                        mode.append(PRE_MULTI)
                        i += 1
                    else:
                        mode.append(ONLY_ONE)
            i += 1
        simCharacter = []
        simMode = []
        i = 0
        while i <= len(character) - 1:
            simCharacter.append(character[i])
            simMode.append(mode[i])
            if mode[i] == PRE_MULTI and i < len(character) - 1 and character[i] == character[i + 1]:
                i += 1
            i += 1
        return simCharacter,simMode


if __name__ == '__main__':
    Solution.isMatch(None, "aaa",
"ab*a*c*a")
