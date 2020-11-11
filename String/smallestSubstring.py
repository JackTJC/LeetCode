import string


class Solution:
    def smallestSubsequence(self, s: str) -> str:
        hashMap = {c: 0 for c in string.ascii_lowercase}
        resList = list()
        for i in range(len(s)):
            hashMap[s[i]] += 1

        return "".join(resList)


if __name__ == '__main__':
    Solution.smallestSubsequence(None, "cbacdcbc")
