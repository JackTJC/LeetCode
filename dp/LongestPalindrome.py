# 5. 最长回文子串
# 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
# python动态规划
class Solution:
    def longestPalindrome(self, s: str) -> str:
        n = len(s)
        if n <= 1:
            return s
        if n < 3:
            return s if s[0] == s[1] else s[0]
        dp = [[0] * n for _ in range(n)]
        maxLen = 1
        ans = s[0]
        # 初始化对角线
        for i in range(n):
            dp[i][i] = 1
        # #初始化对角线上面那条线
        for i in range(n - 1):
            dp[i][i + 1] = 1 if s[i] == s[i + 1] else 0
            if maxLen < 2 and dp[i][i + 1] == 1:
                maxLen = 2
                ans = s[i:i + 2]
        # 计算dp
        num = 2
        while num < n:
            i, j = 0, num
            while j < n:
                dp[i][j] = 1 if (s[j] == s[i] and dp[i + 1][j - 1]) else 0
                if dp[i][j] == 1 and j - i + 1 > maxLen:
                    maxLen = j - i + 1
                    ans = s[i:j + 1]
                j += 1
                i += 1
            num += 1
        return ans

# Java通过
# class Solution {
#     public String longestPalindrome(String s) {
#         boolean [][] dp = new boolean[s.length()][s.length()];
#         String ans="";
#         for(int i=0;i<s.length();i++)
#         {
#             for(int j=0;j<s.length()-i;j++)
#             {
#                 if(i==0)
#                     dp[j][i+j]=true;
#                 else if (i==1)
#                     dp[j][i+j]=(s.charAt(j)==s.charAt(i+j));
#                 else
#                     dp[j][i+j]=(dp[j+1][i+j-1]&&s.charAt(j)==s.charAt(i+j));
#                 if(dp[j][i+j]&&(i+1>ans.length()))
#                     ans=s.substring(j,i+j+1);
#             }
#         }
#         return ans;
#     }
# }
