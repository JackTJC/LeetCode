# 经典问题：求最大公共序列

class Solution:
    def longestCommonSubsequence(self, text1: str, text2: str) -> int:
        # dp
        m = len(text1)
        n = len(text2)
        dp = [[0] * (n + 1) for _ in range(m + 1)]
        for i in range(1, m + 1):
            for j in range(1, n + 1):
                if text1[i - 1] == text2[j - 1]:
                    # 为什么这里是dp[i-1][j-1]
                    # 这里判断条件是text1[i-1]==text2[j-1]，i-1位置和j-1位置的字符都用了一次
                    # 使用dp[i][j-1]和dp[i-1][j]明显造成重复使用
                    dp[i][j] = dp[i - 1][j - 1] + 1
                else:
                    dp[i][j] = max(dp[i][j - 1], dp[i - 1][j])
        return dp[m][n]


if __name__ == '__main__':
    Solution.longestCommonSubsequence(None, "bsbininm", "jmjkbkjkv")
