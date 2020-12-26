# 经典问题：求最大公共子串
class Solution:
    def LCS(self, str1, str2):
        # write code here
        ans = ""
        m = len(str1)
        n = len(str2)
        dp = [[0] * (n + 1) for _ in range(m + 1)]
        for i in range(1, m + 1):
            for j in range(1, n + 1):
                if str1[i - 1] == str2[j - 1]:
                    dp[i][j] = max(dp[i][j - 1], dp[i - 1][j], dp[i - 1][j - 1]) + 1
                else:
                    dp[i][j] = 0
                if dp[i][j] > len(ans):
                    ans = str1[i - dp[i][j]:i + 1]
        return ans if len(ans) > 0 else "-1"


if __name__ == '__main__':
    print(Solution.LCS(None, "xyzcdabc", "xcdabcdef"))
