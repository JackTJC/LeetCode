def leastRevise(s1: str, s2: str) -> int:
    dp = [[0 for j in range(len(s2) + 1)] for i in range(len(s1) + 1)]
    for i in range(len(s1) + 1):
        dp[i][0] = i
    for j in range(len(s2) + 1):
        dp[0][j] = j
    for i in range(1, len(s1) + 1):
        for j in range(1, len(s2) + 1):
            dp[i][j] = min(dp[i - 1][j - 1] + (0 if s1[i-1] == s2[j-1] else 1), dp[i - 1][j] + 1, dp[i][j - 1] + 1)
    return dp[len(s1)][len(s2)]


if __name__ == '__main__':
    print(leastRevise("sfdxbqw", "gfdgw"))
