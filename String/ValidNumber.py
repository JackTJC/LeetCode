class Solution:
    def isNumber(self, s: str) -> bool:
        # State
        AFTER_DOT = 0
        AFTER_E = 1
        AFTER_PLUS = 2
        AFTER_MINUS = 3
        INITIAL = 4
        status = 4
        for i in range(len(s)):
            if status == INITIAL:
                if s[i].isalpha():
                    if s[i] == 'e':
                        if i == len(s) - 1:
                            return False
                        if i == 0:
                            return False
                        else:
                            status = AFTER_E
                    else:
                        return False
                elif s[i] == '.':
                    if i==0 and i==len(s)-1:
                        return False
                    status = AFTER_DOT
                elif s[i] == '-':
                    status = AFTER_MINUS
                elif s[i] == '+':
                    status = AFTER_PLUS
            elif status == AFTER_DOT:
                if not s[i].isdigit():
                    if s[i] == 'e':
                        if i == len(s) - 1:
                            return False
                        status = AFTER_E
                    else:
                        return False
            elif status == AFTER_E:
                if not s[i].isdigit():
                    return False
            elif status == AFTER_MINUS:
                if not s[i].isdigit():
                    if s[i] == 'e' and s[i - 1].isdigit():
                        if i == len(s) - 1:
                            return False
                        status = AFTER_E
                    elif s[i] == '.':
                        status = AFTER_DOT
                    else:
                        return False
            elif status == AFTER_PLUS:
                if not s[i].isdigit():
                    if s[i] == 'e' and s[i - 1].isdigit():
                        if i == len(s) - 1:
                            return False
                        status = AFTER_E
                    elif s[i] == '.':
                        status = AFTER_DOT
                    else:
                        return False
        return True
