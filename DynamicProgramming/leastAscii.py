# 712. 两个字符串的最小ASCII删除和
# 给定两个字符串s1, s2，找到使两个字符串相等所需删除字符的ASCII值的最小和。
class Solution:
    def minimumDeleteSum(self, s1: str, s2: str) -> int:
        m=len(s1)
        n=len(s2)
        dp=[[0]*(n+1) for _ in range(m+1)]
        for i in range(m+1):
            for j in range(n+1):
                if i==0:
                    if j==0:
                        dp[i][j]=0
                    else:
                        dp[i][j]=dp[i][j-1]+ord(s2[j-1])
                elif j==0:
                    dp[i][j]=dp[i-1][j]+ord(s1[i-1])
                else:
                    if s1[i-1]==s2[j-1]:
                        dp[i][j]=dp[i-1][j-1]
                    else:
                        dp[i][j]=min(dp[i-1][j]+ord(s1[i-1]),dp[i][j-1]+ord(s2[j-1]))
        return dp[m][n]